package actions

import (
	"fmt"
	"janic0/gemserv/redis"
	"janic0/gemserv/utils"
	"strings"

	"github.com/cvhariharan/gemini-server"
)

func DeletePage(w *gemini.Response, r *gemini.Request, path string) {
	input := utils.GetInput(r)
	if len(input) == 0 {
		w.SetStatus(gemini.StatusInput, "Enter path to delete:")
		w.SendStatus()
		return
	}
	input = strings.TrimRight(input, "/")
	if input != path {
		w.SetStatus(gemini.StatusInput, "Try again:")
		w.SendStatus()
		return
	}
	err := redis.Client.Del("page:" + path).Err()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to delete page.")
		w.SendStatus()
		return
	}
	w.SetStatus(gemini.StatusRedirect, fmt.Sprintf("/%s/admin", utils.AdminSession))
	w.SendStatus()
}
