package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndentedJSON(c *gin.Context, indentedJsonWithoutErr func(*gin.Context), err error) {
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		indentedJsonWithoutErr(c)
	}
}
