package strava

import (
	"time"
)

// OauthResponse is the return type from DoTokenExchange
type OauthResponse struct {
	TokenType    string  `json:"token_type"`
	ExpiresAt    int64   `json:"expires_at"`
	ExpiresIn    int64   `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
	AccessToken  string  `json:"access_token"`
	Athlete      Athlete `json:"athlete"`
}

// Athlete is
// TODO: Experiement with embedded structs here
type Athlete struct {
	ID                     int64     `json:"id"`
	Username               string    `json:"username"`
	ResourceState          int64     `json:"resource_state"`
	FirstName              string    `json:"firstname"`
	LastName               string    `json:"lastname"`
	City                   string    `json:"city"`
	State                  string    `json:"state"`
	Country                string    `json:"country"`
	Sex                    string    `json:"sex"`
	PremiumSubscriber      bool      `json:"premium"`
	SummitSubscriber       bool      `json:"summit"`
	AccountCreationDate    time.Time `json:"created_at"`
	LastAccountUpdatedDate time.Time `json:"updated_at"`
	BadgeType              int64     `json:"badge_type_id"`
	MediumAvatarURL        string    `json:"profile_medium"`
	LargeAvatarURL         string    `json:"profile"`
	// `json:"friend"` - what is this, I get null
	// `json:"follower"` - what is this, I get null
}

type Avatar struct {
	MediumAvatarURL string `json:"profile_medium"`
	LargeAvatarURL  string `json:"profile"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type Personal struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Sex       string `json:"sex"`
}

type Meta struct {
	ID                     int64     `json:"id"`
	Username               string    `json:"username"`
	PremiumSubscriber      bool      `json:"premium"`
	SummitSubscriber       bool      `json:"summit"`
	AccountCreationDate    time.Time `json:"created_at"`
	LastAccountUpdatedDate time.Time `json:"updated_at"`
	BadgeType              int64     `json:"badge_type_id"`
	ResourceState          int64     `json:"resource_state"`
}

// ResourceState is utility for embedding as most models share it
type ResourceState struct {
	ResourceState int64 `json:"resource_state,omitempty"`
}

type SummaryActivityAthlete struct {
	ID int64 `json:"id"`
	ResourceState
}

type SummaryActivityMap struct {
	ID              string `json:"id"`
	SummaryPolyline string `json:"summary_polyline"`
	Polyline        string `json:"polyline"`
	ResourceState
}

//"timezone": "(GMT-08:00) America/Los_Angeles",
type SummaryActivity struct {
	ResourceState
	Athlete              SummaryActivityAthlete `json:"athlete"`
	Name                 string                 `json:"name"`
	Distance             float64                `json:"distance"`
	MovingTime           int64                  `json:"moving_time"`
	ElapsedTime          int64                  `json:"elapsed_time"`
	TotalElevationGain   float64                `json:"total_elevation_gain"`
	Type                 string                 `json:"type"`
	WorkoutType          int64                  `json:"workout_type"`
	ID                   int64                  `json:"id"`
	ExternalID           string                 `json:"external_id"`
	UploadID             int64                  `json:"upload_id"`
	StartDate            time.Time              `json:"start_date"`
	StartDateLocal       time.Time              `json:"start_date_local"`
	UTCOffset            float64                `json:"utc_offset"`
	StartLatLng          []float64              `json:"start_latlng"`
	EndLatLng            []float64              `json:"end_latlng"`
	LocationCity         string                 `json:"location_city"`
	LocationState        string                 `json:"location_state"`
	LocationCountry      string                 `json:"location_country"`
	AchievementCount     int64                  `json:"achievement_count"`
	KudosCount           int64                  `json:"kudos_count"`
	CommentCount         int64                  `json:"comment_count"`
	AthleteCount         int64                  `json:"athlete_count"`
	PhotoCount           int64                  `json:"photo_count"`
	Map                  SummaryActivityMap     `json:"map"`
	Trainer              bool                   `json:"trainer"`
	Commute              bool                   `json:"commute"`
	Manual               bool                   `json:"manual"`
	Private              bool                   `json:"private"`
	Flagged              bool                   `json:"flagged"`
	GearID               string                 `json:"gear_id"`
	FromAcceptedTag      bool                   `json:"from_accepted_tag"`
	AverageSpped         float64                `json:"average_speed"`
	MaxSpeed             float64                `json:"max_speed"`
	AverageCadence       float64                `json:"average_cadence"`
	AverageWatts         float64                `json:"average_watts"`
	WeigthedAverageWatts int64                  `json:"weighted_average_watts"`
	KiloJoules           float64                `json:"kilojoules"`
	DeviceWatts          bool                   `json:"device_watts"`
	HasHeartrate         bool                   `json:"has_heartrate"`
	AverageHeartrate     float64                `json:"average_heartrate"`
	MaxHeartrate         float64                `json:"max_heartrate"`
	MaxWatts             int64                  `json:"max_watts"`
	PRCount              int64                  `json:"pr_count"`
	TotalPhotoCount      int64                  `json:"total_photo_count"`
	HasKudoed            bool                   `json:"has_kudoed"`
	SufferScore          int64                  `json:"suffer_score"`
}

// TODO: This does not include segment efforts because I don't pay and cannot be bothered to type the model
type DetailedActivity struct {
	SummaryActivity
	MaxElevaton    float64              `json:"elev_high"`
	MinElevation   float64              `json:"elev_low"`
	Timezone       string               `json:"timezone"`
	UploadIDString string               `json:"upload_id_str"`
	Description    string               `json:description"`
	Photos         ActivityPhotoSummary `json:"photos"`
	Gear           ActivityGearSummary  `json:"gear"`
	Calories       float64              `json:"calories"`
	DeviceName     string               `json:"device_name"`
	EmbedToken     string               `json:"embed_token"`
	SplitsMetric   RunningSplit         `json:"splits_metric"`
	SplitsStandard RunningSplit         `json:"splits_standard"`
}

type RunningSplit struct {
	AverageSpeed        float64 `json:"average_speed"`
	Distance            float64 `json:"distance"`
	ElapsedTime         int64   `json:"elapsed_time"`
	ElevationDifference float64 `json:"elevation_difference"`
	PaceZone            int64   `json:"pace_zone"`
	MovingTime          int64   `json:"moving_time"`
	Split               int64   `json:"split"`
}

type ActivityGearSummary struct {
	ResourceState
	ID       string  `json:"id"`
	Primary  bool    `json:"primary"`
	Name     string  `json:"name"`
	Distance float64 `json:"distance"s`
}

type ActivityPhotoSummary struct {
	Count   int64                `json:"count"`
	Primary ActivityPhotoPrimary `json:"primary"`
}

type ActivityPhotoPrimary struct {
	ID       int64  `json:"id"`
	Source   int64  `json:"source"`
	UniqueID string `json:"unique_id"`
	URLs     string `json:"urls"`
}
