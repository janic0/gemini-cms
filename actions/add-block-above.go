package actions

import (
	"janic0/gemserv/generators"
	"janic0/gemserv/redis"
	"janic0/gemserv/utils"

	"github.com/cvhariharan/gemini-server"
	"github.com/google/uuid"
)

func AddBlockAbove(w *gemini.Response, r *gemini.Request, path string, blockIndex int64) {
	input := utils.GetInput(r)
	if len(input) == 0 {
		w.SetStatus(gemini.StatusInput, "Block content:")
		w.SendStatus()
		return
	}
	key := "page:" + path
	value, err := redis.Client.LIndex(key, blockIndex).Result()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Block not found.")
		w.SendStatus()
		return
	}
	pipe := redis.Client.Pipeline()
	changeId := uuid.NewString()
	pipe.LSet(key, blockIndex, changeId).Err()
	pipe.LInsertBefore(key, changeId, input).Err()
	pipe.LSet(key, blockIndex+1, value).Err()
	_, err = pipe.Exec()
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Pipeline failed.")
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
