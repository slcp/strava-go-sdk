package strava_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stuartforrest-infinity/strava"
	"net/http"
	"testing"
)

func TestAuthHeader_Set(t *testing.T) {
	req := http.Request{}
	c := strava.NewClientWithSecret(12345, "test-secret", http.DefaultClient)

	c.SetAuthHeader("test-token", &req)

	assert.Equal(t, http.Header{
		"Authorization": {"Bearer test-token"},
	}, req.Header)
}