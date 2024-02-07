package pkg

import (
	"strings"

	"github.com/google/uuid"
)

func NewRequestID() string {
	id := uuid.New()
	return strings.ReplaceAll(id.String(), "-", "")
}
