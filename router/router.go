package router

import "github.com/gin-gonic/gin"
import . "example/web-service-gin/controller"

func InitRouter() *gin.Engine {
	router := gin.Default()
    router.GET("/albums", GetAlbums)
    router.GET("/albums/:id", GetAlbumByID)
    router.POST("/albums", PostAlbums)
    router.PUT("/albums/:id", UpdateAlbumById)
    router.DELETE("/albums/:id", DeleteAlbumById)

    return router
}