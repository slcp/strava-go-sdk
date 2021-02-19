package strava_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stuartforrest-infinity/strava"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_GetDetailedActivity(t *testing.T) {
	validResponse := makeMockDetailedActivity(int64(1))

	tests := []struct {
		name string
		resp func() *http.Response
		expE error
		expO strava.GetDetailedActivityResponse
	}{
		{
			name: "should return error if response code is greater than 300",
			resp: func() *http.Response {
				return &http.Response{
					StatusCode: 400,
					Body:       ioutil.NopCloser(bytes.NewBufferString("strava api error")),
					// Must be set to non-nil value or it panics
					Header: make(http.Header),
				}
			},
			expE: fmt.Errorf("status code greater than 300 received with message: strava api error"),
		},
		{
			name: "success",
			resp: func() *http.Response {
				or := validResponse
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
			expO: strava.GetDetailedActivityResponse{
				Activity: validResponse,
			},
		},
	}

	for _, tc := range tests {
		client := NewTestClient(func(req *http.Request) *http.Response {
			assert.Equal(t, "www.strava.com", req.URL.Hostname())
			assert.Equal(t, "/api/v3/activities/12345", req.URL.Path)
			return tc.resp()
		})
		c := strava.NewClientWithSecret(12345, "test-secret", client)

		o, e := c.GetDetailedActivity("test-token", strava.GetDetailedActivityData{
			ActivityID: 12345,
		})

		if e != nil {
			assert.Equal(t, tc.expE.Error(), e.Error())
			break
		}
		if tc.expE == nil {
			assert.NoError(t, e)
		}
		assert.Equal(t, tc.expO, o)
	}
}

func makeMockDetailedActivity(i int64) strava.DetailedActivity {
	return strava.DetailedActivity{
		SummaryActivity: strava.SummaryActivity{
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
		},
	}
}
