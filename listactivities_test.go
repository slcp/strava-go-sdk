package strava_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stuartforrest-infinity/strava"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func Test_ListActivities(t *testing.T) {
	validListActivitiesResponse := func() []strava.SummaryActivity {
		var o []strava.SummaryActivity
		for i := 0; i < 3; i++ {
			o = append(o, makeMockActivity(int64(i)))
		}
		return o
	}
	validUrlQueryValues := url.Values{
		"before":   {"12345"},
		"after":    {"54321"},
		"page":     {"10"},
		"per_page": {"100"},
	}

	tests := []struct {
		name string
		resp func() *http.Response
		expE error
		expO strava.ListActivitiesResponse
		expV url.Values
	}{
		{
			name: "strava returns an error status code",
			resp: func() *http.Response {
				return &http.Response{
					StatusCode: 500,
					Body:       ioutil.NopCloser(bytes.NewBufferString("strava api error")),
					// Must be set to non-nil value or it panics
					Header: make(http.Header),
				}
			},
			expV: validUrlQueryValues,
			expE: errors.New("status code greater than 300 received with message: strava api error"),
		},
		{
			name: "success",
			resp: func() *http.Response {
				or := validListActivitiesResponse()
				b, e := json.Marshal(or)
				if e != nil {
					t.Fatal("failed to marshal response body")
				}
				return &http.Response{
					StatusCode: 200,
					// Send response to be tested
					Body: ioutil.NopCloser(bytes.NewBuffer(b)),
					// Must be set to non-nil value or it panics
					Header: make(http.Header),
				}
			},
			expV: validUrlQueryValues,
			expO: strava.ListActivitiesResponse{
				Activities: validListActivitiesResponse(),
			},
		},
	}

	for _, tc := range tests {
		client := NewTestClient(func(req *http.Request) *http.Response {
			assert.Equal(t, "www.strava.com", req.URL.Hostname())
			assert.Equal(t, "/api/v3/athlete/activities", req.URL.Path)
			assert.Equal(t, tc.expV, req.URL.Query())
			return tc.resp()
		})
		c := strava.NewClientWithSecret(12345, "test-secret", client)

		o, e := c.ListActivities("test-token", strava.ListActivitiesData{
			Before:  12345,
			After:   54321,
			Page:    10,
			PerPage: 100,
		})

		if e != nil {
			assert.Equal(t, tc.expE.Error(), e.Error())
			break
		}
		assert.Equal(t, tc.expO, o)
	}
}

func makeMockActivity(i int64) strava.SummaryActivity {
	return strava.SummaryActivity{
		ResourceState: strava.ResourceState{ResourceState: i},
		Athlete: strava.SummaryActivityAthlete{
			ID:            i,
			ResourceState: strava.ResourceState{ResourceState: i},
		},
		Name:               fmt.Sprintf("test-activity-name-%d", i),
		Distance:           float64(i),
		MovingTime:         i,
		ElapsedTime:        i,
		TotalElevationGain: float64(i),
		Type:               "",
		WorkoutType:        i,
		ID:                 i,
		ExternalID:         "",
		UploadID:           i,
		StartDate:          time.Time{},
		StartDateLocal:     time.Time{},
		UTCOffset:          float64(i),
		StartLatLng:        nil,
		EndLatLng:          nil,
		LocationCity:       "",
		LocationState:      "",
		LocationCountry:    "",
		AchievementCount:   i,
		KudosCount:         i,
		CommentCount:       i,
		AthleteCount:       i,
		PhotoCount:         i,
		Map: strava.SummaryActivityMap{
			ID:              "",
			SummaryPolyline: fmt.Sprintf("test-activity-summary-polyline-%d", i),
			Polyline:        fmt.Sprintf("test-activity-polyline-%d", i),
			ResourceState:   strava.ResourceState{ResourceState: i},
		},
		Trainer:              false,
		Commute:              false,
		Manual:               false,
		Private:              false,
		Flagged:              false,
		GearID:               "",
		FromAcceptedTag:      false,
		AverageSpped:         float64(i),
		MaxSpeed:             float64(i),
		AverageCadence:       float64(i),
		AverageWatts:         float64(i),
		WeigthedAverageWatts: i,
		KiloJoules:           float64(i),
		DeviceWatts:          false,
		HasHeartrate:         false,
		AverageHeartrate:     float64(i),
		MaxHeartrate:         float64(i),
		MaxWatts:             i,
		PRCount:              i,
		TotalPhotoCount:      i,
		HasKudoed:            false,
		SufferScore:          i,
	}
}
