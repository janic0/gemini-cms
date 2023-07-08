package actions

import (
	"janic0/gemserv/generators"
	"janic0/gemserv/redis"

	"github.com/cvhariharan/gemini-server"
)

func EnablePage(w *gemini.Response, r *gemini.Request, path string) {
	err := redis.Client.LSet("page:"+path, 0, "enabled").Err()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to enable page.")
		w.SendStatus()
		return
	}
	content, err := generators.GetPageEditContent(path)
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to get edit page.")
		w.SendStatus()
		return
	}
	w.SetStatus(gemini.StatusSuccess, "text/gemini")
	w.Write(content)
}

func DisablePage(w *gemini.Response, r *gemini.Request, path string) {
	err := redis.Client.LSet("page:"+path, 0, "disabled").Err()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to disable page.")
		w.SendStatus()
		return
	}
	content, err := generators.GetPageEditContent(path)
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to get edit page.")
		w.SendStatus()
		return
	}
	w.SetStatus(gemini.StatusSuccess, "text/gemini")
	w.Write(content)
}
