package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
