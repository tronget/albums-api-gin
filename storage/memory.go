package storage

import (
	"errors"
	"fmt"

	"github.com/tronget/albums-rest-api-gin/model"
)

var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums() []model.Album {
	return albums
}

func AddAlbum(newAlbum model.Album) error {
	if isIDExists(newAlbum.ID) {
		return fmt.Errorf("album with ID '%s' already exists", newAlbum.ID)
	}
	albums = append(albums, newAlbum)
	return nil
}

func FindAlbumByID(id string) (*model.Album, error) {
	for i, a := range albums {
		if a.ID == id {
			return &albums[i], nil
		}
	}
	return nil, errors.New("album not found")
}

func isIDExists(id string) bool {
	for i := range albums {
		if albums[i].ID == id {
			return true
		}
	}
	return false
}
