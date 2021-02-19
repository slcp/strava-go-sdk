package strava

import (
	"fmt"
	"strconv"
	"time"
)

type SubscriptionRequestValidationBody struct {
	Challenge string `json:"hub.challenge"`
}

func CreateSubscriptionRequestValidationBody(c string) SubscriptionRequestValidationBody {
	return SubscriptionRequestValidationBody{
		Challenge: c,
	}
}

type WebhookEventAspectType string

const (
	UpdateWebhookEventAspectType WebhookEventAspectType = "update"
	DeleteWebhookEventAspectType WebhookEventAspectType = "delete"
	CreateWebhookEventAspectType WebhookEventAspectType = "create"
)

type WebhookEventBody struct {
	AspectType     WebhookEventAspectType `json:"aspect_type"`
	Time           EventTime              `json:"event_time"`
	ObjectID       int64                  `json:"object_id"`
	ObjectType     string                 `json:"object_type"`
	StravaUserID   int64                  `json:"owner_id"`
	SubscriptionID int64                  `json:"subscription_id"`
	Updates        map[string]interface{} `json:"updates"`
}

// EventTime is time.Time with a custom Unmarshaller
type EventTime time.Time

// Implements UnmarshallJSON interface - expects string representing Unix timestamp
func (t *EventTime) UnmarshalJSON(s []byte) (err error) {
	ts, e := strconv.ParseInt(string(s), 10, 64)
	if e != nil {
		fmt.Println("cannot parse timestamp string as int")
	}
	*(*time.Time)(t) = time.Unix(ts, 0).UTC()
	return
}
