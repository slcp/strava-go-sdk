package strava
import (
	"fmt"
	"net/http"
)


func (c Client) SetAuthHeader(token string, req *http.Request) {
	// Initialise headers if there are none
	if len((*req).Header) == 0 {
		(*req).Header = http.Header{}
	}
	(*req).Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", token)}
	return
}
