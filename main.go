package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type studo struct {
	ID string `json:"id"`
}

var studos = []studo{
	{ID: "0"},
}

func main() {
	router := gin.Default()
	router.GET("/studos", getStudos)
	router.Run("localhost:10030")
}

func getStudos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, studos)
}
