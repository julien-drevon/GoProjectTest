package core

import (
	"encoding/json"
	"errors"
	"os"
)

func SerializeFile[T any](path string, context T) error {
	dataFile, err := os.Create(path)
	if err != nil {
		return err
	}
	ser3, err := json.Marshal(context)
	if err != nil {
		return err
	}
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

func UnserializeFile[T any](path string) (T, error) {
	var context T
	if IsExistFile(path) {
		var context T
		data, err := os.ReadFile(path)
		if err != nil {
			return context, err
		}
		err = json.Unmarshal(data, &context)

		return context, err
	}
	return context, errors.New("File " + path + " not exist")
}
