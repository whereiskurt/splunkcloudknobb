package service

import (
	"encoding/json"
	"fmt"
)

// Report holds and object representing KO Dashboard
type Report struct {
	Name            string
	ID              string
	Author          string
	Updated         string
	Owner           string
	Digest          string
	Search          string
	QualifiedSearch string
	RawEntry        string //json returned from the call that built it.
}

// ListReport returns all the reports for the user
func (s *Service) ListReport(auth AuthCookies, chanr chan interface{}) error {
	defer close(chanr)

	err := s.ConsumePagedResults(auth, chanr, translateReport, auth.URL, "reportResultsURL")
	return err
}

//translateReport converts byte[] to SplunkReportDataUIView and outputs Report onto the channel
func translateReport(body *[]byte, chanr chan interface{}) (total int, count int, err error) {
	var src SplunkReportDataUIView
	err = json.Unmarshal(*body, &src)
	if err != nil {
		err = fmt.Errorf("failed to convert body contents for ListReport to JSON: %v", err)
		return 0, 0, err
	}

	for _, e := range src.Entry {
		var r Report
		r.Name = e.Name
		r.Author = e.Author
		r.Search = e.Content.Search
		r.QualifiedSearch = e.Content.QualifiedSearch
		r.Owner = e.ACL.Owner
		r.Updated = e.Updated

		bb, err := json.MarshalIndent(e, "", "  ")
		if err != nil {
			return 0, 0, err
		}
		r.RawEntry = string(bb)
		chanr <- r
	}

	return src.Paging.Total, len(src.Entry), nil
}
