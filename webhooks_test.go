package strava_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stuartforrest-infinity/strava"
	"testing"
	"time"
)

func TestEventTime_UnmarshalJSON(t *testing.T) {
	var actual strava.EventTime
	expected := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)

	bytes, e := json.Marshal(expected.Unix())
	assert.NoError(t, e, "expected to be able to marshall timestamp")

	e = json.Unmarshal(bytes, &actual)
	assert.NoError(t, e, "unmarshalling")
	assert.Equal(t, expected.String(), time.Time(actual).String())
}
