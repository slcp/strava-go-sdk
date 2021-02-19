package strava

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ListActivitiesData struct {
	Before  int
	After   int
	Page    int
	PerPage int
}

type ListActivitiesResponse struct {
	Activities []SummaryActivity
}

func (c Client) ListActivities(token string, d ListActivitiesData) (ListActivitiesResponse, error) {
	ru := "https://www.strava.com/api/v3/athlete/activities"
	qs := make(map[string]int, 4)
	if d.Before >= 0 {
		qs["before"] = d.Before
	}
	if d.After >= 0 {
		qs["after"] = d.After
	}
	if d.PerPage > 1 {
		qs["page"] = d.Page
	}
	if d.PerPage > 0 {
		qs["per_page"] = d.PerPage
	}
	reqUrl, e := newUrlWithQueryStrings(ru, qs)
	if e != nil {
		return ListActivitiesResponse{}, e
	}

	req := &http.Request{
		Method:           http.MethodGet,
		URL:              reqUrl,
	}
	c.SetAuthHeader(token, req)
	r, e := c.HttpDo.Do(req)
	if e != nil {
		return ListActivitiesResponse{}, e
	}
	defer r.Body.Close()
	if r.StatusCode > 300 {
		b, _ := ioutil.ReadAll(r.Body)
		return ListActivitiesResponse{}, fmt.Errorf("status code greater than 300 received with message: %s", string(b))
	}
	var da []SummaryActivity
	decoder := json.NewDecoder(r.Body)
	e = decoder.Decode(&da)
	if e != nil {
		return ListActivitiesResponse{}, e
	}
	return ListActivitiesResponse{
		Activities: da,
	}, nil
}

func newUrlWithQueryStrings(ru string, query map[string]int) (*url.URL, error) {
	u := fmt.Sprintf("%s?", ru)
	for k, v := range query {
		if v == 0 {
			continue
		}
		u = fmt.Sprintf("%s%s=%d&", u, k, v)
	}

	reqUrl, e := url.Parse(u)
	reqUrl.Query()
	if e != nil {
		return &url.URL{}, e
	}
	return reqUrl, nil
}
