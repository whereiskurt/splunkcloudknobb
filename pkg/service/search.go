package service

import (
	"encoding/json"
	"fmt"
	"log"
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
func (s *Service) ListSearchHistory(auth AuthCookies, chansh chan interface{}) error {

	defer close(chansh)

	spl := s.RenderTemplateOneLine("searchHistorySearchSPL")
	jobSID, err := s.submitSearchJob(auth, spl)
	if err != nil {
		return err
	}
	s.SessionMap["JobSID"] = jobSID

	err = s.waitForStatusDone(auth)
	if err != nil {
		return err
	}

	err = s.ConsumePagedResults(auth, chansh, translateSearchHistory, auth.URL, "searchHistoryResultsURL")

	return err
}

func translateSearchHistory(body *[]byte, chansh chan interface{}) (total int, count int, err error) {

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
