package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func Test_EmptyWhere(t *testing.T) {
	assert := assert.New(t)

	expected := []TypeForTest{}
	actual := Where(
		[]TypeForTest{},
		func(x TypeForTest) bool { return x.Name == "42" })

	assert.Equal(expected, actual)
}
