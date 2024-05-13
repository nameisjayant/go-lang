package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []album{
	{"1", "Blue Train", "John Coltrane", "56.99"},
	{"2", "5 Am Club", "Robin Sharma", "90.88"},
}

func main() {
	route := gin.Default()
	route.GET("/albums", getAlbums)
	route.GET("/albums/:id", getAlbumByID)
	route.POST("/albums", postAlbums)
	route.Run("localhost:8080")
}

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

func getAlbums(c *gin.Context) {
	response := Response{
		Status:  true,
		Message: "Albums Fetch Successfully",
		Data:    albums,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

type Response struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Data    []album `json:"data"`
}
