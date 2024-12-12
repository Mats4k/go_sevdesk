package go_sevdesk

import (
	"fmt"
	"strings"
)

const ApiURL = "https://my.sevdesk.de/api/v1"

func buildURL(endpoint string) string {
	if strings.HasPrefix(endpoint, "/") {
		return fmt.Sprintf("%s%s", ApiURL, endpoint)
	}
	return fmt.Sprintf("%s/%s", ApiURL, endpoint)
}
