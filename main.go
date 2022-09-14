package main

import (
	"file-share/controllers"
	"file-share/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/tracks", controllers.GetAllTracks)

	route.Run()
}
