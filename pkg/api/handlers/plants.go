package plants

import (
	"fmt"
	"net/http"

	"thirsty-jotter/pkg/api/models"
	"thirsty-jotter/pkg/database"

	"github.com/gin-gonic/gin"
	xg "github.com/kerdokurs/xata-go"
)

// plants slice to seed the plants data
var plants = []models.Plant{
	{Id: "1", Name: "papadie", Description: "mica", Type: "flower", Image: ""},
	{Id: "2", Name: "sunflower", Description: "mica", Type: "flower", Image: ""},
	{Id: "3", Name: "Parsley", Description: "mica", Type: "herb", Image: ""},
}

// GetPlants response with the list of plants as JSON
func GetPlants(c *gin.Context) {
	// this serialized the struct into JSON and adds it as response
	c.IndentedJSON(http.StatusOK, plants)
}

// GetXataPlants retrieves a list of plants from xata db
func GetXataPlants(c *gin.Context) {

	// query from xata
	xata := database.BuildXataClient()

	fmt.Println("Connecting to xata...")
	query := xata.Plants.Select("*", "plants.*").Filter("quantity", xg.Gt, 0)

	// this serialized the struct into JSON and adds it as response
	c.IndentedJSON(http.StatusOK, query)
}

// PostPlant adds a plant received as a payload as JSON in the request body
func PostPlant(c *gin.Context) {
	var newPlant models.Plant

	// c.BindJSON binds the request body to newPlant var
	if err := c.BindJSON(&newPlant); err != nil {
		return
	}

	plants = append(plants, newPlant)
	c.IndentedJSON(http.StatusCreated, plants)

}

// GetPlantByID retrieve the plant which id matches the plant id passed as a parameter to the function
// then returns that plant as a response
func GetPlantByID(c *gin.Context) {
	// TODO: revise this logic
	// id := c.Param("id")

	// for _, p := range plants {
	// 	if p.ID == id {
	// 		c.IndentedJSON(http.StatusOK, p)
	// 		return
	// 	}

	// }
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "plant not found"})
}
