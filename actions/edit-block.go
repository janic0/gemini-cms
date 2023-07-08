package actions

import (
	"fmt"
	"janic0/gemserv/generators"
	"janic0/gemserv/redis"
	"janic0/gemserv/utils"

	"github.com/cvhariharan/gemini-server"
)

func EditBlock(w *gemini.Response, r *gemini.Request, path string, block int64) {
	input := utils.GetInput(r)
	if len(input) == 0 {
		w.SetStatus(gemini.StatusInput, "Block content:")
		w.SendStatus()
		return
	}
	err := redis.Client.LSet("page:"+path, block, input).Err()
	if err != nil {
		fmt.Print(err.Error())
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to edit block")
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
