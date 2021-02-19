package strava_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stuartforrest-infinity/strava"
	"net/http"
	"testing"
)

func TestClient_NewClient(t *testing.T) {
	id := 1234
	secret := "secret"

	c := strava.NewClientWithSecret(id, secret, http.DefaultClient)

	assert.Equal(t, strava.Client{
		ID:     id,
		Secret: secret,
		HttpDo: &http.Client{},
	}, c)
}
