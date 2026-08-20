package pkg

import (
	"github.com/google/uuid"
)

func ParseUUID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)

	if err != nil {
		return uuid.UUID{}, err
	}

	return parsedID, nil
}
