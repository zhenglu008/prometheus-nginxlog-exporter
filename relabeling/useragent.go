package relabeling

import (
	"github.com/mssola/user_agent"
)

func Parse(userAgent string) []string {
	var result []string
	ua := user_agent.New(userAgent)
	result = append(result, ua.Model())
	return result
}