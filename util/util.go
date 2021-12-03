package util

import (
	"github.com/google/uuid"
)

func UUIDv4() string {
	return uuid.NewString()
}
