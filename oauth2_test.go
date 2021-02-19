package strava_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stuartforrest-infinity/strava"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)


func TestOauth2_DoTokenExchange(t *testing.T) {
	validOauthResponse := strava.OauthResponse{
		TokenType:    "Bearer",
		ExpiresAt:    time.Date(2020, 01, 01, 01, 01, 01, 01, time.UTC).Unix(),
		ExpiresIn:    100,
		RefreshToken: "refresh-token",
		AccessToken:  "access-token",
		Athlete: strava.Athlete{
			ID: 12345,
		},
	}

	tests := []struct {
		name string
		resp func() *http.Response
		expE error
		expO strava.OauthResponse
	}{
		{
			name: "strava returns an error status code",
			resp: func() *http.Response {
				return &http.Response{
					StatusCode: 500,
					Body: ioutil.NopCloser(bytes.NewBufferString("strava api error")),
					// Must be set to non-nil value or it panics
					Header: make(http.Header),
				}
			},
			expE: errors.New("status code greater than 300 received with message: strava api error"),
		},
		{
			name: "success",
			resp: func() *http.Response {
				or := validOauthResponse
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
			expO: validOauthResponse,
		},
	}

	for _, tc := range tests {
		client := NewTestClient(func(req *http.Request) *http.Response {
			assert.Equal(t, req.URL.String(), "https://www.strava.com/api/v3/oauth/token")
			return tc.resp()
		})
		c := strava.NewClientWithSecret(12345, "test-secret", client)

		o, e := c.DoTokenExchange("test-code")

		if e != nil {
			assert.Equal(t, tc.expE.Error(), e.Error())
			break
		}
		assert.Equal(t, tc.expO, o)
	}
}
