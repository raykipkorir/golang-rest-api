package main

import (
	// "fmt"
	"net/http"
	// "log"
	"strconv"
	"github.com/gin-gonic/gin"
)

type album struct{
	ID 		string		`json:"id"`
	Title 	string 		`json:"title"`
	Artist 	string 		`json:"artist"`
	Price 	float64  	`json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Destined 2 win", Artist: "Lil Tjay", Price: 100},
	{ID: "2", Title: "Forest hills drive", Artist: "J Cole", Price: 100},
	{ID: "3", Title: "Shoot for the stars Aim for the moon", Artist: "Pop Smoke", Price: 100},
	{ID: "4", Title: "Legends never die", Artist: "Juice Wrld", Price: 100},
}

func main(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// res := map[string]any{"message": "pong"}
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/albums", getAlbums)
	r.POST("/albums", createAblum)
	r.GET("/albums/:id", getAlbumByID)
	r.DELETE("/albums/:id", deleteAlbum)
	r.Run() // listen and serve on 0.0.0.0:8080
}

// gin.Context is the most part of gin. It carries request details, validates and serializes JSON, and many more

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums) // serializes the struct into JSON and adds it to the response.
}

func createAblum(c *gin.Context){
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil{
		return 
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context){
	id := c.Param("id")
	for _, album := range albums{
		if album.ID == id{
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func deleteAlbum(c *gin.Context){
	id := c.Param("id")
	for _, album := range albums{
		if album.ID == id{
			// convert str to int
			x, _ := strconv.Atoi(id)
			remove(albums, x-1)
			c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Deleted successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// remove an element from a slice
func remove(slice []album, index int){
	albums = append(slice[:index], slice[index+1:]...)
}
