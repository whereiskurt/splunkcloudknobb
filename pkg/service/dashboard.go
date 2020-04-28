package service

import (
	"encoding/json"
	"fmt"
	"log"
	urls "net/url"

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
func (s *Service) ListDashboard(auth AuthCookies, chand chan Dashboard) (err error) {
	defer close(chand)

	const query = `output_mode=json&search=((isDashboard=1 AND ((rootNode="dashboard" AND version=1) OR rootNode="form" OR rootNode="view" OR rootNode="html") AND isVisible=1) AND ((eai:acl.sharing="user" AND eai:acl.owner="%s") OR (eai:acl.sharing!="user")))&sort_dir=asc&sort_key=label&sort_mode=alpha&sort_mode=alpha&count=%d&offset=%d`
	const size = 100

	offset := 0

PAGING:
	for {
		u := auth.URL + `splunkd/__raw/servicesNS/-/search/data/ui/views?`
		url := u + urls.PathEscape(fmt.Sprintf(query, auth.Username, size, offset))

		client, req, err := s.authGetRequest(url)
		if err != nil {
			err := fmt.Errorf("failed to create request to ListDashboards")
			return err
		}
		s.authCookieDecorate(auth, client, req)

		body, err := s.retryRequest("Dashboard", client, req)
		if err != nil {
			return err
		}

		total, count, err := translateDashboard(&body, chand)
		if err != nil {
			return err
		}

		offset = offset + count
		if total <= offset || size != count {
			break PAGING
		}

	}

	return nil
}

func translateDashboard(body *[]byte, chand chan Dashboard) (total int, count int, err error) {
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
