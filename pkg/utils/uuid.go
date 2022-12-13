package utils

import (
	"fmt"

	"github.com/gofrs/uuid"

	"github.com/rinatkh/test_2022/internal/constants"
)

func GenUUID() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("%w: %v", constants.ErrGenerateUUID, err)
	}
	return uuid.String(), nil
}
