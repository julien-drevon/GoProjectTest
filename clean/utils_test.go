package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsExistFile_ShouyldBeTrue(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(IsExistFile("utils.go"), true)
}

func Test_IsExistFile_ShouyldBeFalse(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(IsExistFile("untruc.go"), false)
}

type SerializeTest struct {
	Test string
}

func Test_Serialize_File(t *testing.T) {
	assert := assert.New(t)
	path := "Test_Serialize_File.json"
	DeleteUtilsTestFile(path)
	SerializeFile(path, SerializeTest{Test: "42"})

	assert.Equal(true, IsExistFile(path))

	file, err := os.ReadFile(path)
	assert.Equal(nil, err)
	assert.Equal("{\"Test\":\"42\"}", string(file))

	DeleteUtilsTestFile(path)
}

func DeleteUtilsTestFile(path string) {

	if IsExistFile(path) {
		os.Remove(path)
	}
}

func Test_Unserialize_File(t *testing.T) {
	assert := assert.New(t)
	path := "Test_Unserialize_File.json"
	DeleteUtilsTestFile(path)
	os.WriteFile(path, []byte("{\"Test\":\"42\"}"), os.ModePerm)

	actual, _ := UnserializeFile[SerializeTest](path)
	assert.Equal(SerializeTest{Test: "42"}, actual)

	DeleteUtilsTestFile(path)
}
