package linq

import (
	"encoding/json"
	"errors"
)

type MappingError struct {
	Message string
}

func Where[T any](typeForTest []T, where func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range typeForTest {
		if where(v) {
			result = append(result, v)
		}
	}
	return result
}

func Select[T any, K any](typeForTest []T, projection func(x T) K) []K {
	result := make([]K, len(typeForTest))
	for i, v := range typeForTest {
		result[i] = projection(v)
	}

	return result
}
func Any[T any](typeForTest []T, any func(T) bool) bool {

	for _, v := range typeForTest {
		if any(v) {
			return true
		}
	}
	return false
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
