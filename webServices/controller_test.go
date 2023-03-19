package main

import (
	//"clean/core"
	// "pokedex/domain"
	// "presenter"

	"testing"

	"github.com/stretchr/testify/assert"
)

// func GetDebile() core.PaginationResult[domain.Pokemon] {
// 	return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
// }

type MyTestStruct struct {
	Id int
}

func TestCreateQuerySimple(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, false)
}
