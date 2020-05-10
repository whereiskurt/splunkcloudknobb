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
