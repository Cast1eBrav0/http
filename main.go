package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

type studo struct {
	ID string `json:"id"`
}

var studos = []studo{
	{ID: "0"},
}

func main() {
	ConfigRuntime()
	StartGin()
}

func getStudos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, studos)
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartGin() {
	router := gin.Default()
	router.GET("/stud", getStudos)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
