package relabeling

import (
	"github.com/mssola/user_agent"
)

func Parse(userAgent string) []string {
	var result []string
	ua := user_agent.New(userAgent)
	engine, _ := ua.Engine()
	result = append(result, ua.OS())
	result = append(result, ua.Platform())
	result = append(result, engine)
	return result
}