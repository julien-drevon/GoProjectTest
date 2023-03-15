package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginationResult(t *testing.T) {
	assert := assert.New(t)
	result := PaginationResult[string]{}
	zeroValue := []string(nil)
	assert.Equal(zeroValue, result.Result)
}

func TestPaginationResult_New_AssignVlues(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1}, 0, 1, 2)
	want := []int{1}
	assert.Equal(want, result.Result)
}

func TestPaginationResult_New_TotalValue(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1}, 42, 1, 1)
	want := 42
	assert.Equal(want, result.Pagination.Total)
}

func TestPaginationResult_New_TotalPage(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1, 1}, 9, 1, 2)
	want := 5
	assert.Equal(want, result.Pagination.TotalPage)
}

func TestPaginationResult_New_CurrentPage(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1, 1}, 9, 2, 2)
	want := 2
	assert.Equal(want, result.Pagination.CurrentPage)
}

func TestPaginationResult_New_CurrentPageShouldBeMajorThanZero(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1, 1}, 9, 0, 2)
	want := 1
	assert.Equal(want, result.Pagination.CurrentPage)
}

func TestPaginationResult_New_HasNext_whenIsTrue(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1, 1}, 9, 2, 2)
	want := true
	assert.Equal(want, result.Pagination.HasNext)
}

func TestPaginationResult_New_HasPrevious_whenIsTrue(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1, 1}, 9, 2, 2)
	want := true
	assert.Equal(want, result.Pagination.HasPrevious)
}

func TestPaginationResult_New_HasPageAsk_IsSet(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{1, 1}, 9, 2, 2)
	want := 2
	assert.Equal(want, result.Pagination.PageSize)
}

func TestPaginationResult_New_EnterBadValue(t *testing.T) {
	assert := assert.New(t)
	result := NewPaginationResult([]int{}, 9, 0, 0)
	want := Pagination{
		CurrentPage: 1,
		HasNext:     true,
		HasPrevious: false,
		TotalPage:   0,
		PageSize:    0,
		Total:       9,
	}
	assert.Equal(want, result.Pagination)
}
