package views

import (
	"janic0/gemserv/redis"
	"janic0/gemserv/utils"
	"strings"

	"github.com/cvhariharan/gemini-server"
)

func AdminDashboard(w *gemini.Response, r *gemini.Request) {
	messages := make([]string, 0)
	messages = append(messages, "# Welcome back.", "", "## Pages", "", utils.CraftAdminLink("/admin/add-page", "+ Page"), "")
	prefix := "page:"
	keys, err := redis.Client.Keys(prefix + "*").Result()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, err.Error())
		w.SendStatus()
		return
	}
	for _, key := range keys {
		path := key[len(prefix):]
		messages = append(messages, utils.CraftAdminLink("/admin/edit-page"+utils.EscapePathSegments(path), path))
	}
	w.Write([]byte(strings.Join(messages, "\r\n")))
}
