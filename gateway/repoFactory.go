package gateway

import (
	"encoding/gob"
	"os"
	"pokedex/domain"
)

func NewRepoForUnitTests() Repo {
	return Repo{Context: make(map[string][]domain.Pokemon), PersistFile: true}
}

func NewRepoForUnitTestsWithContext(context map[string][]domain.Pokemon) Repo {
	return Repo{Context: context, PersistFile: true}
}

func NewRepoForWithPersistance() Repo {
	var data map[string][]domain.Pokemon
	dataFile, err := openPersistanceFile()
	DeserializePersistanceFile(err, data, dataFile)

	return Repo{Context: data, PersistFile: true}
}

func DeserializePersistanceFile(err error, data map[string][]domain.Pokemon, dataFile *os.File) {
	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data)
	if err != nil {
		os.Exit(1)
	}
	dataFile.Close()
}

func openPersistanceFile() (*os.File, error) {
	dataFile, err := os.Open(pathOfPersistanceFile)
	if err != nil {
		os.Exit(1)
	}

	return dataFile, err
}
