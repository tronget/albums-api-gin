package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tronget/albums-rest-api-gin/model"
	"github.com/tronget/albums-rest-api-gin/storage"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.GetAlbums())
}

func PostAlbum(c *gin.Context) {
	newAlbum := model.Album{}
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := storage.AddAlbum(newAlbum); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := storage.FindAlbumByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
