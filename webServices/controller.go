package webservices

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Controller example
type Controller struct {
}

// NewController example
func NewController() *Controller {
	return &Controller{}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func ConvertToQuery[T any](initialValue T, conversion func(T) (T, error)) (T, error) {
	val, err := conversion(initialValue)
	if err != nil {
		return val, NewHttpError(http.StatusBadRequest, errors.New(""))
	}
	return val, err
}

func MultiConvertToQuery[T any](initialValue T, conversions []func(T) (T, error)) (T, error) {

	preVal := initialValue
	for _, conversion := range conversions {
		val, _ := ConvertToQuery(preVal, conversion)
		preVal = val
	}

	return preVal, nil
}

// HTTPError example

func NewHttpError(status int, err error) error {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	//ctx.JSON(status, er)
	json, _ := json.Marshal(er)
	return errors.New(string(json))
}
