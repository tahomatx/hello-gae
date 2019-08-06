package hello

import (
	"github.com/satori/go.uuid"
)

func Hi() string {
	// return "heireaoijoi"
    return uuid.NewV4().String()
}
