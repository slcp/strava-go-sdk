package strava

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type GetDetailedActivityData struct {
	ActivityID int64
}

type GetDetailedActivityResponse struct {
	Activity DetailedActivity
}

func (c Client) GetDetailedActivity(token string, d GetDetailedActivityData) (GetDetailedActivityResponse, error) {
	ru := "https://www.strava.com/api/v3/activities"
	reqUrl := fmt.Sprintf("%s/%d", ru, d.ActivityID)

	r, e := c.HttpDo.Get(reqUrl)
	if e != nil {
		return GetDetailedActivityResponse{}, e
	}

	defer r.Body.Close()
	if r.StatusCode > 300 {
		b, _ := ioutil.ReadAll(r.Body)
		return GetDetailedActivityResponse{}, fmt.Errorf("status code greater than 300 received with message: %s", string(b))
	}
	var da DetailedActivity
	decoder := json.NewDecoder(r.Body)
	e = decoder.Decode(&da)
	if e != nil {
		return GetDetailedActivityResponse{}, e
	}
	return GetDetailedActivityResponse{
		Activity: da,
	}, nil
}
