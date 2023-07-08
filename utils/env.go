package utils

import (
	"os"

	"github.com/google/uuid"
)

var Password = os.Getenv("ADMIN_PASSWORD")
var AdminSession = uuid.NewString()
var LastAdminSession = uuid.NewString()
