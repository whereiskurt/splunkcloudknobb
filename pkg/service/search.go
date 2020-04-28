package service

import (
	"encoding/json"
	"fmt"
	"log"
	urls "net/url"
	"strings"
	"time"
)

// SearchHistoryResult are returned for each search history entry
type SearchHistoryResult struct {
	SearchID     string
	Search       string
	ResultCount  string
	TotalRunTime string
	Time         string
	RawEntry     string //json returned from the call that built it.
}

// ListSearchHistory fills the channel with search history entry
func (s *Service) ListSearchHistory(auth AuthCookies, chansh chan SearchHistoryResult) (err error) {
	defer close(chansh)

	search := urls.QueryEscape(`| history | search NOT search="| history*" AND NOT search="*metadata*" AND NOT search="*loadjob*" AND NOT savedsearch_name="*" AND NOT search="search" AND NOT search="*from sid*" AND NOT search="| eventcount summarize=false index=* index=_**" AND NOT provenance="UI:LocateData" AND NOT provenance="UI:TableEditor" AND NOT provenance="UI:DataModel" AND NOT provenance="UI:Pivot" AND NOT provenance="UI:Dataset" | dedup search | head 100000`)

	jobSID, err := s.submitSearchJob(auth, search)
	if err != nil {
		return err
	}

	err = s.waitForDone(auth, jobSID)
	if err != nil {
		return err
	}

	const size = 100
	const query = `splunkd/__raw/servicesNS/nobody/search/search/jobs/%s/results?output_mode=json&offset=%d&count=%d&search=%s`

	offset := 0
PAGING:
	for {
		// Non of these params need ay escaping.
		url := auth.URL + fmt.Sprintf(query, jobSID, offset, size, urls.PathEscape("|search"))

		client, req, err := s.authGetRequest(url)
		if err != nil {
			err := fmt.Errorf("failed to create request for Search History Results")
			return err
		}
		s.authCookieDecorate(auth, client, req)

		body, err := s.retryRequest("SearchHistoryResults", client, req)
		if err != nil {
			return err
		}

		total, count, err := translateSearchHistory(&body, chansh)
		if err != nil {
			return err
		}

		offset = offset + count
		if total <= offset || size != count {
			break PAGING
		}

	}

	return
}

func translateSearchHistory(body *[]byte, chansh chan SearchHistoryResult) (total int, count int, err error) {

	var src SplunkSearchResultsUIView
	err = json.Unmarshal(*body, &src)
	if err != nil {
		err = fmt.Errorf("failed to convert body contents for Search History Results from JSON: %v", err)
		return 0, 0, err
	}

	for _, e := range src.Results {
		var se SearchHistoryResult
		se.SearchID = e.Sid
		se.Search = e.Search
		se.TotalRunTime = e.TotalRunTime
		se.Time = e.Time

		bb, err := json.MarshalIndent(e, "", "  ")
		if err != nil {
			log.Fatalf("JSON marshal unexpected error")
		}
		se.RawEntry = string(bb)
		chansh <- se
	}
	total = src.PostProcessCount
	count = len(src.Results)

	return total, count, nil
}

func (s *Service) submitSearchJob(auth AuthCookies, search string) (sid string, err error) {

	searchBody := fmt.Sprintf(`rf=*&auto_cancel=30&status_buckets=300&output_mode=json&search=%s&earliest_time=0&preview=false&provenance=UI:Search`, search)

	jobURL := auth.URL + fmt.Sprintf(`splunkd/__raw/servicesNS/%s/search/search/jobs`, urls.PathEscape(auth.Username))

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
	jobSID, err := translateJobSID(&searchbody)
	if err != nil {
		return "", err
	}

	return jobSID, nil
}

func translateJobSID(body *[]byte) (sid string, err error) {
	var src SplunkJobSubmission
	err = json.Unmarshal(*body, &src)
	if err != nil {
		err = fmt.Errorf("failed to extract jobSID from submission response: %v", err)
		return "", err
	}

	return src.SID, nil
}

func (s *Service) waitForDone(auth AuthCookies, jobSID string) (reterr error) {
	const size = 100
	const maxloops = 15
	var loopcount = 0

CHECK:
	for {
		// 3) Sleep for 1 seconds and let the query finishing running
		time.Sleep(2 * time.Second)

		const statusURL = `splunkd/__raw/servicesNS/%s/search/search/jobs/%s?output_mode=json`
		url := auth.URL + fmt.Sprintf(statusURL, urls.PathEscape(auth.Username), urls.PathEscape(jobSID))

		client, req, err := s.authGetRequest(url)
		if err != nil {
			err := fmt.Errorf("failed to create request to check search history job status: %v", err)
			return err
		}
		s.authCookieDecorate(auth, client, req)

		body, err := s.retryRequest("SearchHistoryStatus", client, req)
		if err != nil {
			return err
		}

		status, err := searchStatus(&body)
		if err != nil {
			return err
		}

		switch strings.ToUpper(status) {
		case "DONE":
			reterr = nil
			break CHECK
		case "FAILED":
			reterr = fmt.Errorf("search history job %s failed with status %s", jobSID, status)
			break CHECK
		default:
			loopcount = loopcount + 1
		}

		if loopcount > maxloops {
			reterr = fmt.Errorf("exceeded waiting retries for job search history")
			break CHECK
		}

	}

	return reterr
}

func searchStatus(body *[]byte) (status string, err error) {
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
