package service

import (
	"encoding/json"
	"errors"
	"fmt"
	urls "net/url"
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
func ListReport(auth AuthCookies, chanr chan Report) error {
	defer close(chanr)

	const query = `output_mode=json&sort_dir=asc&sort_key=name&sort_mode=natural&sort_mode=natural&search=NOT ((is_scheduled=1 AND (alert_type!=always OR alert.track=1 OR (dispatch.earliest_time="rt*" AND dispatch.latest_time="rt*" AND actions="*" AND actions!="")))) AND (eai:acl.owner="%s") AND ((eai:acl.sharing="user" AND eai:acl.owner="%s") OR (eai:acl.sharing!="user")) AND is_visible=1&count=%d&offset=%d`
	const size = 100

	offset := 0
PAGING:
	for {
		u := auth.URL + `splunkd/__raw/servicesNS/-/search/saved/searches?`
		url := u + urls.PathEscape(fmt.Sprintf(query, auth.Username, auth.Username, size, offset))

		client, req, err := authGetRequest(url)
		if err != nil {
			err := errors.New("failed to create request to ListReports")
			return err
		}
		authCookieDecorate(auth, client, req)
		body, err := retryRequest("Report", client, req)
		if err != nil {
			return err
		}

		total, count, err := translateReport(&body, chanr)
		if err != nil {
			return err
		} else if count <= 0 {
			return fmt.Errorf("no results returned in ListReports")
		}

		offset = offset + count
		if total <= offset || size != count {
			break PAGING
		}
	}

	return nil
}

//translateReport converts byte[] to SplunkReportDataUIView and outputs Report onto the channel
func translateReport(body *[]byte, chanr chan Report) (total int, count int, err error) {
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
