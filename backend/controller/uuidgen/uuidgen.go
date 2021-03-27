package uuidgen

import (
	"github.com/google/uuid"
)

// UUIDGenerator is interface to generate uuid.
type UUIDGenerator interface {
	GenNewRandom() (string, error)
}

// uuidGenerator has generating method.
type uuidGenerator struct {
}

func NewUUIDGenerator() UUIDGenerator {
	return &uuidGenerator{}
}

// GenNewRandom wrap uuid.NewRandom()
func (g *uuidGenerator) GenNewRandom() (string, error) {
	random, err := uuid.NewRandom()
	return random.String(), err
}
