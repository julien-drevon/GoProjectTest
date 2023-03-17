package linq

import (
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
