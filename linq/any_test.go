package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
