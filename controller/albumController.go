package controller

import "github.com/gin-gonic/gin"
import "net/http"
import . "example/web-service-gin/data"
import . "example/web-service-gin/model"

func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, Albums)
}

func PostAlbums(c *gin.Context) {
    var newAlbum Album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    Albums = append(Albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")
    
    for _, a := range Albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func DeleteAlbumById(c *gin.Context) {
    id := c.Param("id")

    for index, a := range Albums {
        if a.ID == id {
            Albums = append(Albums[:index], Albums[index+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func UpdateAlbumById(c *gin.Context) {
    id := c.Param("id")

    for index, a := range Albums {
        if a.ID == id {
            if err := c.BindJSON(&a); err != nil {
                return
            }
            Albums[index] = a
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}