package actions

import (
	"janic0/gemserv/generators"
	"janic0/gemserv/redis"
	"janic0/gemserv/utils"
	"strings"

	"github.com/cvhariharan/gemini-server"
)

func AddPage(w *gemini.Response, r *gemini.Request) {
	input := utils.GetInput(r)
	if input == "" || !strings.HasPrefix(input, "/") {
		w.SetStatus(gemini.StatusInput, "Enter page path:")
		w.SendStatus()
		return
	}
	input = strings.TrimRight(input, "/")
	if input == "" {
		input = "/"
	}
	err := redis.Client.RPush("page:"+input, "disabled", "# "+input).Err()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to create page.")
		w.SendStatus()
		return
	}
	editContent, err := generators.GetPageEditContent(input)
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to get edit page.")
		w.SendStatus()
	} else {
		w.Write(editContent)
	}
}
