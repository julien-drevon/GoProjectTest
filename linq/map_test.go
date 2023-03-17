package linq

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapperSimple(t *testing.T) {
	assert := assert.New(t)

	query := MyTestStruct{}
	expect := MyTestStruct{Names: []string{"a"}}
	actual, _ := Mapper(query, func(q MyTestStruct) (MyTestStruct, error) {
		q.Names = []string{"a"}
		return q, nil
	})

	assert.Equal(expect, actual)
}

func TestMultiMapper(t *testing.T) {
	assert := assert.New(t)

	query := MyTestStruct{}
	expect := MyTestStruct{Names: []string{"a1"}}
	actual, _ := MultiMapper(query, []func(MyTestStruct) (MyTestStruct, error){
		func(q MyTestStruct) (MyTestStruct, error) {
			q.Names = []string{"a"}
			return q, nil
		},
		func(q MyTestStruct) (MyTestStruct, error) {
			q.Names = Select(q.Names, func(x string) string { return x + "1" })
			return q, nil
		}})

	assert.Equal(expect, actual)
}
func TestMapperWithError(t *testing.T) {
	assert := assert.New(t)

	err := "{\"Message\":\"\"}"
	expect := string(err)
	_, actual := Mapper(MyTestStruct{}, func(q MyTestStruct) (MyTestStruct, error) {
		id, internalErr := strconv.Atoi("a")
		q.Id = id
		return q, internalErr
	})

	assert.Equal(expect, actual.Error())
}

func TestMultiMapperError(t *testing.T) {
	assert := assert.New(t)

	query := MyTestStruct{}
	expect := NewMappingError("")
	ret, actual := MultiMapper(query, []func(MyTestStruct) (MyTestStruct, error){
		func(q MyTestStruct) (MyTestStruct, error) {
			q.Names = []string{"a"}
			return q, errors.New((""))
		},
		func(q MyTestStruct) (MyTestStruct, error) {
			q.Names = Select(q.Names, func(x string) string { return x + "1" })
			return q, nil
		}})
	assert.Equal(MyTestStruct{Names: []string{"a"}}, ret)
	assert.Equal(expect, actual)
}
