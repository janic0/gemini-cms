package views

import (
	"janic0/gemserv/generators"

	"github.com/cvhariharan/gemini-server"
)

func EditPage(w *gemini.Response, r *gemini.Request, page string) {
	content, err := generators.GetPageEditContent(page)
	if err != nil {
		w.SetStatus(gemini.StatusTemporaryFailure, "Failed to get edit page.")
		w.SendStatus()
		return
	}
	w.Write(content)
}
