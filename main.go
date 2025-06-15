package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tronget/albums-rest-api-gin/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumById)
	router.POST("/albums", handlers.PostAlbum)

	router.Run("localhost:8080")
}
