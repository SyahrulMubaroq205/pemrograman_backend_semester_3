package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "blue", Price: 56.99},
	{ID: "2", Title: "orange", Price: 76.89},
	{ID: "3", Title: "red", Price: 12.02},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumsId(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})	
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsId)
	router.PUT("/albums/:id", updateAlbumsId)
	router.DELETE("/albums/:id", destroyAlbumsId)

	router.Run("localhost:8000")
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if error := c.BindJSON(&newAlbum); error != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func updateAlbumsId(c *gin.Context) {
	id := c.Param("id")
	var updateAlbum album

	if error := c.BindJSON(&updateAlbum); error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	// Find and update the album
	for i, album := range albums {
		if album.ID == id {
			albums[i].Title = updateAlbum.Title
			albums[i].Price = updateAlbum.Price
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func destroyAlbumsId(c *gin.Context) {
	id := c.Param("id")

	// Find and destroy the album
	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}