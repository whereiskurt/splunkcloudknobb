package service

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// LookupFile holds and object representing KO Dashboard
type LookupFile struct {
	Filename string
	Path     string
	Owner    string
	App      string
	Sharing  string
	Status   string
}

// ListLookupFiles returns all the reports for the user
func (s *Service) ListLookupFiles(auth AuthCookies, chanr chan interface{}) error {
	defer close(chanr)

	err := s.ConsumePagedResults(auth, chanr, s.translateLookupFilesList, auth.URL, "lookfilesResultsURL")
	return err
}

// RemoveLookupFile will try and remove the lookup file
func (s *Service) RemoveLookupFile(auth AuthCookies) error {
	resultsURL := s.RenderTemplateOneLine("lookupFileRemoveURL")

	url := auth.URL + resultsURL

	body := s.RenderTemplateOneLine("lookupFileRemovePostBody")
	body = strings.ReplaceAll(body, "/", "%2F")

	client, req, err := s.authPostRequest(url, strings.NewReader(body))
	if err != nil {
		return err
	}
	s.authCookieDecorate(auth, client, req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	ret, code, err := s.retryRequest("LookupFileRemove", client, req)
	if err != nil {
		return err
	}

	if len(ret) == 0 {
		return fmt.Errorf("fatal: failed to delete - empty return")
	}

	s.Log.Debugf("The return code was: %d", code)

	return nil
}

func (s *Service) RestoreLookupFile(auth AuthCookies, filename, path, namespace string) {

	// curl 'https://es-cgleis.splunkcloud.com/en-US/manager/search/data/lookup-table-files/_new' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Content-Type: multipart/form-data; boundary=---------------------------108658978517118608851849409995' -H 'Connection: keep-alive' -H 'Cookie: splunkweb_csrf_token_8443=5421638032078838654; token_key=5421638032078838654; experience_id=b8e2c2af-037d-df81-e24f-25823ef3e699; splunkweb_uid=E822EE85-BE1B-4B1D-AC11-8BEF57662817; login=true; session_id_8443=44c4d5743e6d0e7566d5e12b6b2fb355f8f0d83a; splunkd_8443=jpggBHpxNyZCuewbV4hWx1UzlInkkIO4MMrnwb6BEkwIt8eEGos_rvA8kLmyjs7ftvwFXh1sUgJhpGVLp52Rv1Y9khFkQokqJdZziSIBNDJ5Ia8lUqVBRfQMncuty2fVzJziOEuz0t9yzkvgrz_FPuGkNVzQ^kf3' -H 'Upgrade-Insecure-Requests: 1' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache' --data-binary $'-----------------------------108658978517118608851849409995\r\nContent-Disposition: form-data; name="__action"\r\n\r\nnew\r\n-----------------------------108658978517118608851849409995\r\nContent-Disposition: form-data; name="__redirect"\r\n\r\n1\r\n-----------------------------108658978517118608851849409995\r\nContent-Disposition: form-data; name="__ns"\r\n\r\nsearch\r\n-----------------------------108658978517118608851849409995\r\nContent-Disposition: form-data; name="spl-ctrl_lookupfile"; filename="test.simple.csv"\r\nContent-Type: text/csv\r\n\r\n-----------------------------108658978517118608851849409995\r\nContent-Disposition: form-data; name="name"\r\n\r\ntesttesttest.csv\r\n-----------------------------108658978517118608851849409995\r\nContent-Disposition: form-data; name="splunk_form_key"\r\n\r\n5421638032078838654\r\n-----------------------------108658978517118608851849409995--\r\n'

	// -----------------------------108658978517118608851849409995
	// Content-Disposition: form-data; name="__action"

	// new
	// -----------------------------108658978517118608851849409995
	// Content-Disposition: form-data; name="__redirect"

	// 1
	// -----------------------------108658978517118608851849409995
	// Content-Disposition: form-data; name="__ns"

	// search
	// -----------------------------108658978517118608851849409995
	// Content-Disposition: form-data; name="spl-ctrl_lookupfile"; filename="test.simple.csv"
	// Content-Type: text/csv

	// header1, header2, header3
	// row1col1, row1col2, row1col3
	// row2col1, row2col2, row2col3
	// row3col1, row3col2, row3col3
	// -----------------------------108658978517118608851849409995
	// Content-Disposition: form-data; name="name"

	// testtesttest.csv
	// -----------------------------108658978517118608851849409995
	// Content-Disposition: form-data; name="splunk_form_key"

	// 5421638032078838654
	// -----------------------------108658978517118608851849409995--

}

//translateLookupFilesList converts byte[] of html to a LookupFile and pump into channel
func (s *Service) translateLookupFilesList(body *[]byte, chanr chan interface{}) (total int, count int, err error) {
	var headings []string

	html := string(*body)
	myRegex, _ := regexp.Compile(`of (\d+) items`)
	found := myRegex.FindString(html)

	if len(found) == 0 {
		return 0, 0, fmt.Errorf("fatal: couldn't find record count in html")
	}

	//Split the string, and the middle item is the ### of total records
	total, err = strconv.Atoi(strings.Split(found, " ")[1])
	if err != nil {
		return 0, 0, fmt.Errorf("fatal: couldn't convert value in '%s' to integer", found)
	}

	r := bytes.NewReader(*body)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return 0, 0, fmt.Errorf("fatal: couldn't parase html for lookup file listing")
	}

	// Use the golang jQuery style library
	var tc = 0
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		// We only want the first table it finds.
		if tc > 0 {
			return
		}
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			var row []string
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
				t := strings.TrimSpace(tableheading.Text())
				t = strings.ReplaceAll(t, "\n", "")
				t = strings.ReplaceAll(t, "\t", "")
				headings = append(headings, t)
			})
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				t := strings.TrimSpace(tablecell.Text())
				t = strings.ReplaceAll(t, "\n", "")
				t = strings.ReplaceAll(t, "\t", "")
				row = append(row, t)
			})
			if len(row) > 4 {
				//TODO: Add some header lookup code incase fields are 're-ordered' or in different orders.
				row[3] = strings.TrimSpace(strings.ReplaceAll(row[3], "| Permissions", ""))

				chanr <- LookupFile{Path: row[0], Filename: filepath.Base(row[0]), Owner: row[1], App: row[2], Sharing: row[3], Status: row[4]}
				count = count + 1
			}
		})
		tc = tc + 1
	})

	return total, count, nil
}
