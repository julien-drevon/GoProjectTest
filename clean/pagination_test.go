package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPaginationResult(t *testing.T) {
// 	assert := assert.New(t)
// 	result := PaginationResult[string]{}

//		assert.Equal(result, nil)
//	}

func TestPaginationResult(t *testing.T) {
	assert := assert.New(t)
	result := PaginationResult[string]{}
	zeroValue := []string(nil)
	assert.Equal(zeroValue, result.Result)
}

func TestPaginationResult_New_AssignVlues(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1}, 0)
	want := []int{1}
	assert.Equal(want, result.Result)
}

func TestPaginationResult_New_TotalValue(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1}, 42)
	want := 42
	assert.Equal(want, result.Pagination.Total)
}

func TestPaginationResult_New_TotalPage(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1, 1}, 9)
	want := 5
	assert.Equal(want, result.Pagination.TotalPage)
}
