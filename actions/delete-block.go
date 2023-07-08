package actions

import (
	"janic0/gemserv/generators"
	"janic0/gemserv/redis"

	"github.com/cvhariharan/gemini-server"
	"github.com/google/uuid"
)

func DelteBlock(w *gemini.Response, r *gemini.Request, path string, blockIndex int64) {
	key := "page:" + path
	changeId := uuid.NewString()
	pipe := redis.Client.Pipeline()
	pipe.LSet(key, blockIndex, changeId)
	pipe.LRem(key, 1, changeId)
	_, err := pipe.Exec()
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
