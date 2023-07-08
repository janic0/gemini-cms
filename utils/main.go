package utils

import (
	"net/url"
	"strings"

	"github.com/cvhariharan/gemini-server"
)

func GetInput(r *gemini.Request) string {
	if len(r.URL.RawQuery) > 0 {
		response, err := url.QueryUnescape(r.URL.RawQuery)
		if err != nil {
			return ""
		} else {
			return response
		}
	} else {
		return ""
	}
}

func CraftAdminLink(path string, text string) string {
	link := "=> "
	link += "/" + AdminSession + path
	if text != "" {
		link += " " + text
	}
	return link
}

func EscapePathSegments(path string) string {
	segments := strings.Split(path, "/")
	for i, v := range segments {
		segments[i] = url.PathEscape(v)
	}
	return strings.Join(segments, "/")
}
