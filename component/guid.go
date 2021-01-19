package component

import (
	"github.com/google/uuid"
)

type GUID struct {
	guid string
}

// NewRandomGUID creates new GUID component with random id value
func NewRandomGUID() GUID {
	return GUID{uuid.New().String()}
}

// NewGUID cretes new GUID component with provided id value
func NewGUID(guid string) GUID {
	return GUID{}
}
