package uuid_utils

import (
	"errors"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func ParseStringToUUID(id string) (uuid.UUID, error) {
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return idUUID, errors.New(err.Error())
	}
	return idUUID, nil
}
