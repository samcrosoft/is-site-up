package entry

import (
	"fmt"
	"net/http"
	"time"
)

// SiteResponse ... The structure of the response to be returned
type SiteResponse struct {
	Status     string
	StatusCode int
}

// CheckSite ... ->something blah blah blha
func CheckSite(siteURL string, timeout time.Duration) (result bool, response *SiteResponse) {

	fmt.Println(fmt.Sprintf("checking site for %s", siteURL))

	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(siteURL)
	if err != nil || resp == nil {
		result = false
	} else {
		result = true
		response = &SiteResponse{Status: resp.Status, StatusCode: resp.StatusCode}
	}
	return
}
