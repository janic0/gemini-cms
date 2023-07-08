package generators

import (
	"fmt"
	"janic0/gemserv/redis"
	"strings"
)

func GetPage(path string) ([]byte, error) {
	messages := make([]string, 0)
	items, err := redis.Client.LRange("page:"+path, 0, -1).Result()
	// escapedPath := utils.EscapePathSegments(path)
	if err != nil {
		return nil, err
	} else if len(items) == 0 {
		return nil, fmt.Errorf("No blocks found.")
	}
	for i, item := range items {
		if i == 0 {
			if item != "enabled" {
				return nil, fmt.Errorf("Page not enabled.")
			}
		} else {
			messages = append(messages, item)
		}
	}
	return []byte(strings.Join(messages, "\r\n")), nil
}
