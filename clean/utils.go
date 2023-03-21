package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func SerializeFile[T any](path string, context T) error {
	dataFile, err := os.Create(path)
	if err != nil {
		return err
	}
	ser3, _ := json.Marshal(context)
	dataFile.Write(ser3)
	if err != nil {
		return err
	}
	return dataFile.Close()
}

func IsExistFile(path string) bool {
	fs, _ := os.Stat(path)
	return fs != nil && !fs.IsDir()
}

func UnserialyzeFile[T any](path string) (T, error) {
	var context T
	if IsExistFile(path) {
		var context T
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return context, err
		}
		err = json.Unmarshal(data, &context)

		return context, err
	}
	return context, errors.New("File" + path + " not exist")
}
