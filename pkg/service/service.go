package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/whereiskurt/splunkcloudknobb/pkg/tmpl"

	"net/http"
	"net/http/cookiejar"
	urls "net/url"
	"time"

	"gopkg.in/matryer/try.v1"
)

// Service holds a transaport and log
type Service struct {
	Log             *log.Logger
	Transport       *http.Transport
	TmplRender      *tmpl.UITemplate
	SessionMap      map[string]string
	ResultCountSize int
}

//NewService creates Service instance
func NewService(log *log.Logger) (s *Service) {
	s = new(Service)
	s.Log = log
	s.Transport = &http.Transport{
		MaxIdleConns:        20,
		IdleConnTimeout:     10 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	s.TmplRender = new(tmpl.UITemplate)

	s.TmplRender.RegisterPackageFile("service/service.tmpl")
	s.TmplRender.RegisterPackageFile("service/search.tmpl")
	s.TmplRender.RegisterPackageFile("service/report.tmpl")

	s.SessionMap = make(map[string]string)

	s.ResultCountSize = 100
	s.SessionMap["Count"] = fmt.Sprintf("%d", s.ResultCountSize)

	return
}

// RenderTemplateOneLine renders tmplname and strips \n \r from
func (s *Service) RenderTemplateOneLine(tmplname string) string {
	rendered := s.TmplRender.RenderPackage(tmplname, s.SessionMap)
	rendered = strings.ReplaceAll(rendered, "\n", "")
	rendered = strings.ReplaceAll(rendered, "\r", "")
	return rendered
}

func (s *Service) decorateUniversalHeaders(req *http.Request) {
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	return
}

func pathEscape(fullurl string) string {
	qm := strings.Index(fullurl, "?")
	if qm < 0 {
		return fullurl
	}

	escaped := fullurl[:qm+1] + urls.PathEscape(fullurl[qm+1:])

	return escaped
}

func (s *Service) authGetRequest(url string) (client *http.Client, req *http.Request, err error) {

	s.Log.Errorf("Before Escaped:%s", url)
	url = pathEscape(url)
	s.Log.Errorf("After Escaped:%s", url)

	client = &http.Client{
		Transport: s.Transport,
		Timeout:   30 * time.Second,
	}

	client.Jar, _ = cookiejar.New(nil)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	s.decorateUniversalHeaders(req)

	return client, req, err
}

func (s *Service) authPostRequest(url string, data io.Reader) (client *http.Client, req *http.Request, err error) {

	url = pathEscape(url)

	client = &http.Client{
		Transport: s.Transport,
		Timeout:   30 * time.Second,
	}

	client.Jar, _ = cookiejar.New(nil)

	req, err = http.NewRequest("POST", url, data)
	if err != nil {
		return nil, nil, err
	}

	s.decorateUniversalHeaders(req)

	return client, req, err
}

var retryIntervals = []int{0, 500, 500, 500, 500, 1000, 1000, 1000, 1000, 1000, 3000}

func (s *Service) sleepBeforeRetry(attempt int) (shouldReRun bool) {
	if attempt < len(retryIntervals) {
		time.Sleep(time.Duration(retryIntervals[attempt]) * time.Millisecond)
		shouldReRun = true
	}
	return
}

func (s *Service) retryRequest(label string, client *http.Client, req *http.Request) ([]byte, error) {
	var body []byte

	err := try.Do(func(attempt int) (bool, error) {
		resp, err1 := client.Do(req)
		if err1 != nil {
			log.Println(fmt.Printf("Failed here: %v, %s", err1, err1))
			return s.sleepBeforeRetry(attempt), err1
		}
		defer resp.Body.Close()

		//TODO: Make this WAY WAY more robust.
		if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
			err2 := fmt.Errorf("failed to get %s: http_resp: %d", label, resp.StatusCode)
			return s.sleepBeforeRetry(attempt), err2
		}

		respBody := resp.Body

		var err2 error
		body, err2 = ioutil.ReadAll(respBody)
		if err2 != nil {
			err2 = fmt.Errorf("failed to read body contents for %s: %v", label, err2)
			return s.sleepBeforeRetry(attempt), err2
		}

		return false, nil
	})

	return body, err
}

func (s *Service) submitSearchJob(auth AuthCookies, spl string) (sid string, err error) {
	spl = urls.QueryEscape(spl)
	s.SessionMap["SearchParam"] = spl

	searchHistoryURL := s.RenderTemplateOneLine("searchHistoryURL")
	jobURL := auth.URL + searchHistoryURL

	searchBody := s.RenderTemplateOneLine("searchHistoryBody")

	client, req, err := s.authPostRequest(jobURL, strings.NewReader(searchBody))
	if err != nil {
		return "", err
	}
	s.authCookieDecorate(auth, client, req)
	req.Header.Add("X-Splunk-Form-Key", auth.SplunkWebCSRF)

	searchbody, err := s.retryRequest("SearchHistoryJob", client, req)
	if err != nil {
		return "", err
	}

	// 2) Unmarshall SplunkJobSubmission the $SID
	jobSID, err := extractJobSID(&searchbody)
	if err != nil {
		return "", err
	}

	s.SessionMap["JobSID"] = jobSID
	return jobSID, nil
}
func extractJobSID(body *[]byte) (sid string, err error) {
	var src SplunkJobSubmission
	err = json.Unmarshal(*body, &src)
	if err != nil {
		err = fmt.Errorf("failed to extract jobSID from submission response: %v", err)
		return "", err
	}

	return src.SID, nil
}

func (s *Service) waitForDone(auth AuthCookies, jobSID string) (reterr error) {

	const maxloops = 15
	var loopcount = 0

CHECK:
	for {
		// 3) Sleep for 1 seconds and let the query finishing running
		time.Sleep(2 * time.Second)

		statusURL := s.RenderTemplateOneLine("searchJobStatusURL")
		url := auth.URL + statusURL

		client, req, err := s.authGetRequest(url)
		if err != nil {
			err := fmt.Errorf("failed to create request to check search history job status: %v", err)
			return err
		}
		s.authCookieDecorate(auth, client, req)

		body, err := s.retryRequest("searchJobStatusURL", client, req)
		if err != nil {
			return err
		}

		status, err := extractSearchStatus(&body)
		if err != nil {
			return err
		}

		switch strings.ToUpper(status) {
		case "DONE":
			reterr = nil
			break CHECK
		case "FAILED":
			reterr = fmt.Errorf("%s job %s failed with status %s", "searchJobStatusURL", jobSID, status)
			break CHECK
		default:
			loopcount = loopcount + 1
		}

		if loopcount > maxloops {
			reterr = fmt.Errorf("exceeded waiting retries for job %s", "searchJobStatusURL")
			break CHECK
		}

	}

	return reterr
}
func extractSearchStatus(body *[]byte) (status string, err error) {
	var src SplunkSearchStatusUIView
	err = json.Unmarshal(*body, &src)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal search status from submission response: %v", err)
		return "", err
	}

	if len(src.Entry) == 0 {
		err = fmt.Errorf("failed to extract search status from submission response: %v", err)
		return "", err
	}

	status = src.Entry[0].Content.DispatchState

	return status, nil
}

// Translator is called after results have return from Paged call
type Translator func(*[]byte, chan interface{}) (int, int, error)

// ConsumePagedResults loops over all the results from baseURL+`tmplname` calling tfn and populated chansh
func (s *Service) ConsumePagedResults(auth AuthCookies, chansh chan interface{}, tfn Translator, baseURL string, tmplname string) error {
	var offset = 0

PAGING:
	for {
		s.SessionMap["Offset"] = fmt.Sprintf("%d", offset)

		resultsURL := s.RenderTemplateOneLine(tmplname)

		url := baseURL + resultsURL

		client, req, err := s.authGetRequest(url)
		if err != nil {
			err := fmt.Errorf("failed to create request for %s", tmplname)
			return err
		}
		s.authCookieDecorate(auth, client, req)

		body, err := s.retryRequest(tmplname, client, req)
		if err != nil {
			log.Fatalf("%v", err)
		}

		total, count, err := tfn(&body, chansh)
		if err != nil {
			return err
		}

		offset = offset + count
		if total <= offset || s.ResultCountSize != count {
			break PAGING
		}
	}

	return nil
}
