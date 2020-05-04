package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// Dashboard holds and object representing KO Dashboard
type Dashboard struct {
	Name     string
	ID       string
	Author   string
	Updated  string
	Owner    string
	Content  string
	RawEntry string //json returned from the call that built it.
}

// ListDashboard returns ALL dashboards returns an array of dasboards from the call
func (s *Service) ListDashboard(auth AuthCookies, chand chan interface{}) error {
	defer close(chand)

	return s.ConsumePagedResults(auth, chand, translateDashboard, auth.URL, "dashboardResultsURL")
}

func translateDashboard(body *[]byte, chand chan interface{}) (total int, count int, err error) {
	var src SplunkDashboardDataUIView
	err = json.Unmarshal(*body, &src)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed to convert body contents for ListDashboard to JSON: %v", err))
		return 0, 0, err
	}

	for _, e := range src.Entry {
		var d Dashboard
		d.Author = e.Author
		d.Owner = e.ACL.Owner
		d.Updated = e.Updated
		d.ID = e.ID
		d.Name = e.Name
		d.Content = e.Content.EaiData
		bb, err := json.MarshalIndent(e, "", "  ")
		if err != nil {
			log.Fatalf("JSON marshal unexpected error")
		}
		d.RawEntry = string(bb)

		chand <- d
	}

	return src.Paging.Total, len(src.Entry), nil
}
