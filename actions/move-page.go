package actions

import (
	"fmt"
	"janic0/gemserv/redis"
	"janic0/gemserv/utils"
	"strings"

	"github.com/cvhariharan/gemini-server"
)

func MovePage(w *gemini.Response, r *gemini.Request, path string) {
	input := utils.GetInput(r)
	if len(input) == 0 {
		w.SetStatus(gemini.StatusInput, "New path:")
		w.SendStatus()
		return
	}
	input = strings.TrimRight(input, "/")
	value, err := redis.Client.RenameNX("page:"+path, "page:"+input).Result()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to move page.")
		w.SendStatus()
		return
	}
	if !value {
		w.SetStatus(gemini.StatusTemporaryFailure, "Page already exists.")
		w.SendStatus()
		return
	}
	w.SetStatus(gemini.StatusRedirect, fmt.Sprintf("/%s/admin/edit-page%s", utils.AdminSession, input))
	w.SendStatus()
}
