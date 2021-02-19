package strava

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type StravaScope string

const (
	StravaScopePrivateActivities      StravaScope = "activity:read_all"
	StravaScopePublicWithoutAcitivies StravaScope = "read"
	StravaScopePrivateithoutAcitivies StravaScope = "read_all"
	StravaScopePrivateProfile         StravaScope = "profile:read_all"
	StravaScopePublicActivities       StravaScope = "activity:read"
)

type RefreshTokensInput struct {
	RefreshToken string
}

func (c Client) DoTokenExchange(code string) (OauthResponse, error) {
	reqUrl := "https://www.strava.com/api/v3/oauth/token"
	grant := "authorization_code"

	r, e := c.HttpDo.PostForm(reqUrl,
		url.Values{"client_id": {fmt.Sprintf("%d", c.ID)}, "client_secret": {c.Secret}, "code": {code}, "grant": {grant}})
	if e != nil {
		return OauthResponse{}, e
	}
	defer r.Body.Close()
	if r.StatusCode > 300 {
		b, _ := ioutil.ReadAll(r.Body)
		return OauthResponse{}, fmt.Errorf("status code greater than 300 received with message: %s", string(b))
	}
	var or OauthResponse
	decoder := json.NewDecoder(r.Body)
	e = decoder.Decode(&or)
	if e != nil {
		return OauthResponse{}, e
	}
	return or, nil
}

func (c Client) RefreshTokens(d RefreshTokensInput) (or OauthResponse, e error) {
	reqUrl := "https://www.strava.com/api/v3/oauth/token"
	grant := "refresh_token"

	r, e := c.HttpDo.PostForm(reqUrl,
		url.Values{"client_id": {fmt.Sprintf("%d", c.ID)}, "client_secret": {c.Secret}, "refresh_token": {d.RefreshToken}, "grant": {grant}})
	if e != nil {
		return
	}
	defer r.Body.Close()
	if r.StatusCode > 300 {
		b, _ := ioutil.ReadAll(r.Body)
		e = fmt.Errorf("status code greater than 300 received with message: %s", string(b))
		return
	}
	decoder := json.NewDecoder(r.Body)
	e = decoder.Decode(&or)
	if e != nil {
		return
	}
	return or, nil
}

func (c Client) GetAuthoriseURL(redirectUrl string, scopes []StravaScope) string {
	stravaBaseUrl := "https://www.strava.com/api/v3/oauth/authorize"
	//apiBaseUrl := "https://k50jbma5l7.execute-api.eu-west-2.amazonaws.com/dev"
	//apiPath := "/strava/auth/callbac"
	//redirectUrl := fmt.Sprintf("%s%s", apiBaseUrl, apiPath)
	//scopes := []StravaScope{StravaScopePrivateithoutAcitivies, StravaScopePrivateActivities}
	return fmt.Sprintf("%s?client_id=%d&response_type=code&scope=%s&redirect_uri=%s", stravaBaseUrl, c.ID, mergeScopes(scopes), redirectUrl)
	//return "https://www.strava.com/api/v3/oauth/authorize?client_id=2350&redirect_uri=https://k50jbma5l7.execute-api.eu-west-2.amazonaws.com/dev/strava/auth/callback&response_type=code&scope=read_all,activity:read_all"
}

func mergeScopes(scopes []StravaScope) string {
	if len(scopes) == 0 {
		return ""
	}
	var (
		sep = []byte(",")
		// preallocate for len(sep) + assume at least 1 character
		out = make([]byte, 0, (1+len(sep))*len(scopes))
	)
	for _, s := range scopes {
		out = append(out, s...)
		out = append(out, sep...)
	}
	return string(out[:len(out)-len(sep)])
}
