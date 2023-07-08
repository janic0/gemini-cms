package actions

import (
	"janic0/gemserv/generators"
	"janic0/gemserv/redis"
	"janic0/gemserv/utils"

	"github.com/cvhariharan/gemini-server"
)

func AddBlock(w *gemini.Response, r *gemini.Request, path string) {
	input := utils.GetInput(r)
	if len(input) == 0 {
		w.SetStatus(gemini.StatusInput, "Block content:")
		w.SendStatus()
		return
	}
	err := redis.Client.RPush("page:"+path, input).Err()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to add block")
		w.SendStatus()
		return
	}
	content, err := generators.GetPageEditContent(path)
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed get page content")
		w.SendStatus()
		return
	}
	w.Write(content)
}
