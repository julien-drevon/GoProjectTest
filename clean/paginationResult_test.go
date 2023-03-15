package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginationResult(t *testing.T) {
	assert := assert.New(t)
	result := PaginationResult[string]{}

	assert.Equal(result, nil)
}
