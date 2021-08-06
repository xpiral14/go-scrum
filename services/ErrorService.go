package services

import (
	"encoding/json"
	"go-scrum/constants"
	"go-scrum/models"
	"strconv"
)

type ErrorService struct{}

func NewErrorService() *ErrorService {
	return new(ErrorService)
}

func (e *ErrorService) ToBytes(error models.Error) []byte {
	err := strconv.Itoa(constants.INVALID_PAYLOAD)

	errorJSON, err2 := json.Marshal(error)
	if err2 != nil {
		return nil
	}
	result := []byte(err)
	return append(result, errorJSON...)
}
