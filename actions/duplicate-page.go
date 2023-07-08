package actions

import (
	"fmt"
	"janic0/gemserv/redis"
	"janic0/gemserv/utils"
	"strings"

	"github.com/cvhariharan/gemini-server"
)

func DuplicatePage(w *gemini.Response, r *gemini.Request, path string) {
	input := utils.GetInput(r)
	if len(input) == 0 {
		w.SetStatus(gemini.StatusInput, "Path to duplicate to:")
		w.SendStatus()
		return
	}
	input = strings.TrimRight(input, "/")
	value, err := redis.Client.LRange("page:"+path, 0, -1).Result()
	fmt.Print("page:" + path)
	fmt.Print(value)
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to read page.")
		w.SendStatus()
		return
	}
	err = redis.Client.RPush("page:"+input, value).Err()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to push new items.")
		w.SendStatus()
		return
	}
	w.SetStatus(gemini.StatusRedirect, fmt.Sprintf("/%s/admin/edit-page%s", utils.AdminSession, utils.EscapePathSegments(input)))
	w.SendStatus()
}
