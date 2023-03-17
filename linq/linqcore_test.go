package linq

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TypeForTest struct {
	Name string
}

func Test_SimpleWhere(t *testing.T) {
	assert := assert.New(t)

	expected := []TypeForTest{{Name: "42"}}
	actual := Where([]TypeForTest{{Name: "42"}, {Name: "moi"}}, (func(x TypeForTest) bool { return x.Name == "42" }))

	assert.Equal(expected, actual)
}
func Test_LongWhere(t *testing.T) {
	assert := assert.New(t)

	expected := []TypeForTest{{Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}}
	actual := Where(
		[]TypeForTest{{Name: "42"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}},
		func(x TypeForTest) bool { return x.Name == "42" })

	assert.Equal(expected, actual)
}

func Test_SimpleSelect(t *testing.T) {
	assert := assert.New(t)

	expected := []string{"42", "moi"}
	actual := Select([]TypeForTest{{Name: "42"}, {Name: "moi"}}, (func(x TypeForTest) string { return x.Name }))

	assert.Equal(expected, actual)
}
func Test_longSelect(t *testing.T) {
	assert := assert.New(t)

	expected := []string{"42", "moi", "moi", "moi", "moi", "moi", "moi", "moi", "moi", "moi", "moi", "moi", "moi", "moi", "moi", "moi", "42", "42", "42", "42", "42", "42", "42", "42", "42"}
	actual := Select(
		[]TypeForTest{{Name: "42"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "moi"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}, {Name: "42"}},
		(func(x TypeForTest) string { return x.Name }))

	assert.Equal(expected, actual)
}
func Test_EmptySelect(t *testing.T) {
	assert := assert.New(t)

	expected := []string{}
	actual := Select(
		[]TypeForTest{},
		(func(x TypeForTest) string { return x.Name }))

	assert.Equal(expected, actual)
}
func Test_EmptyWhere(t *testing.T) {
	assert := assert.New(t)

	expected := []TypeForTest{}
	actual := Where(
		[]TypeForTest{},
		func(x TypeForTest) bool { return x.Name == "42" })

	assert.Equal(expected, actual)
}

func Test_AnyTrue(t *testing.T) {
	assert := assert.New(t)

	expected := true
	actual := Any(
		[]TypeForTest{{Name: "42"}, {Name: "moi"}},
		func(x TypeForTest) bool { return x.Name == "moi" })

	assert.Equal(expected, actual)
}

func Test_AnyFalse(t *testing.T) {
	assert := assert.New(t)

	expected := false
	actual := Any(
		[]TypeForTest{{Name: "42"}, {Name: "toi"}},
		func(x TypeForTest) bool { return x.Name == "moi" })

	assert.Equal(expected, actual)
}

type MyTestStruct struct {
	Id    int
	Names []string
}

func TestCreateQuerySimple(t *testing.T) {
	assert := assert.New(t)

	query := MyTestStruct{}
	expect := MyTestStruct{Names: []string{"a"}}
	actual, _ := Mapper(query, func(q MyTestStruct) (MyTestStruct, error) {
		q.Names = []string{"a"}
		return q, nil
	})

	assert.Equal(expect, actual)
}

func TestCreateMultiQuery(t *testing.T) {
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
func TestCreateQuerySimpleWithError(t *testing.T) {
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

func TestCreateMultiQueryError(t *testing.T) {
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
