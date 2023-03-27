package webServicesPackage

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReferential(t *testing.T) {
	router := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/referential", nil)
	if err != nil {
		t.Fatal(err)
	}

	Run().ServeHTTP(router, req)

	assert.Equal(t, http.StatusOK, router.Code)
	assert.Equal(t, "{\n    \"Result\": [\n        {\n            \"name\": \"pikatchu\"\n        },\n        {\n            \"name\": \"tortank\"\n        },\n        {\n            \"name\": \"draco feu\"\n        },\n        {\n            \"name\": \"bulbizare\"\n        }\n    ],\n    \"Pagination\": {\n        \"CurrentPage\": 1,\n        \"hasNext\": false,\n        \"hasPrevious\": false,\n        \"totalPage\": 1,\n        \"pageSize\": 0,\n        \"total\": 4\n    }\n}", router.Body.String())
}

func TestPostPokemon(t *testing.T) {
	router := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/post", nil)
	if err != nil {
		t.Fatal(err)
	}

	Run().ServeHTTP(router, req)

	assert.Equal(t, http.StatusCreated, router.Code)
	assert.Equal(t, "{\"name\":\"Test Pokemon\",\"type\":\"grass\"}\n", router.Body.String())
}
