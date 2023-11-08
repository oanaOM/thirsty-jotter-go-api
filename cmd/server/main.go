package main

import (
	"fmt"
	plants "thirsty-jotter/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := serveApplication()

	r.Run("localhost:8080")
	fmt.Println("Server running on port 8080")
}

func serveApplication() *gin.Engine {
	router := gin.Default()
	router.GET("/plants", plants.GetPlants)
	router.GET("/xata_plants", plants.GetXataPlants)
	router.GET("/plants/:id", plants.GetPlantByID)
	router.POST("/plants", plants.PostPlant)

	return router

}
