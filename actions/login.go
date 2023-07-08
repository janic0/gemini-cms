package actions

import (
	"fmt"
	"janic0/gemserv/utils"

	"github.com/cvhariharan/gemini-server"
	"github.com/google/uuid"
)

func Login(w *gemini.Response, r *gemini.Request) {
	input := utils.GetInput(r)
	if input != utils.Password {
		w.SetStatus(11, "Password:")
		w.SendStatus()
	} else {
		utils.AdminSession = uuid.NewString()
		w.SetStatus(gemini.StatusRedirect, fmt.Sprintf("/%s/admin", utils.AdminSession))
		w.SendStatus()
	}
}
