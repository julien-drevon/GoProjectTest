package linq

import (
	"encoding/json"
	"errors"
)

type MappingError struct {
	Message string
}

func Mapper[T any](initialValue T, conversion func(T) (T, error)) (T, error) {
	val, err := conversion(initialValue)
	if err != nil {
		return val, NewMappingError("")
	}

	return val, err
}

func MultiMapper[T any](initialValue T, conversions []func(T) (T, error)) (T, error) {
	preVal := initialValue
	var err error
	for _, conversion := range conversions {
		preVal, err = Mapper(preVal, conversion)
		if err != nil {
			return preVal, err
		}
	}

	return preVal, nil
}

func NewMappingError(err string) error {
	er := MappingError{
		Message: err,
	}

	jsonByte, _ := json.Marshal(er)
	return errors.New(string(jsonByte))
}
