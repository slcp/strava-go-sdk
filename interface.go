package strava

type IClient interface {
	DoTokenExchange(code string) (OauthResponse, error)
	ListActivities(token string, d ListActivitiesData) (ListActivitiesResponse, error)
	RefreshTokens(d RefreshTokensInput) (or OauthResponse, e error)
	GetDetailedActivity(token string, d GetDetailedActivityData) (GetDetailedActivityResponse, error)
}
