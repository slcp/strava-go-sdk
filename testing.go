package strava

import (
	"fmt"
	"time"
)

func MakeMockDetailedActivity(i int64) DetailedActivity {
	return DetailedActivity{
		SummaryActivity: MakeMockActivity(i),
		MaxElevaton:     float64(i),
		MinElevation:    float64(i),
		Timezone:        "",
		UploadIDString:  "",
		Description:     "",
		Photos:          ActivityPhotoSummary{
			Count:   i,
			Primary: ActivityPhotoPrimary{
				ID:       i,
				Source:   i,
				UniqueID: "",
				URLs:     "",
			},
		},
		Gear:            ActivityGearSummary{
			ResourceState: ResourceState{ResourceState: i},
			ID:            "",
			Primary:       false,
			Name:          "",
			Distance:      float64(i),
		},
		Calories:        float64(i),
		DeviceName:      "",
		EmbedToken:      "",
		SplitsMetric:    RunningSplit{
			AverageSpeed:        float64(i),
			Distance:            float64(i),
			ElapsedTime:         i,
			ElevationDifference: float64(i),
			PaceZone:            i,
			MovingTime:          i,
			Split:               i,
		},
		SplitsStandard:    RunningSplit{
			AverageSpeed:        float64(i),
			Distance:            float64(i),
			ElapsedTime:         i,
			ElevationDifference: float64(i),
			PaceZone:            i,
			MovingTime:          i,
			Split:               i,
		},
	}
}
func MakeMockActivity(i int64) SummaryActivity {
	return SummaryActivity{
		ResourceState: ResourceState{ResourceState: i},
		Athlete: SummaryActivityAthlete{
			ID:            i,
			ResourceState: ResourceState{ResourceState: i},
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
		Map: SummaryActivityMap{
			ID:              "",
			SummaryPolyline: fmt.Sprintf("test-activity-summary-polyline-%d", i),
			Polyline:        fmt.Sprintf("test-activity-polyline-%d", i),
			ResourceState:   ResourceState{ResourceState: i},
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
