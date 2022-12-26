package main

import (
	"io"
	"net/http"
	"os"

	"github.com/MelvinKim/golang-gin-gonic/controller"
	"github.com/MelvinKim/golang-gin-gonic/middlewares"
	"github.com/MelvinKim/golang-gin-gonic/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	// create a log file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	// load static assets eg CSS files
	server.Static("/css", "./templates/css")

	// load HTML templates
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())
	// gindump.Dump() middleware --> adds more metadata about the request or response
	// middlewares.Logger() --> logging middleware
	// server.Use(gin.Logger())

	// API grouping
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video input is valid",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", nil)
	}

	server.Run(":8080")
}

// package main

// import (
// 	"github.com/MelvinKim/golang-gin-gonic/controller"
// 	"github.com/MelvinKim/golang-gin-gonic/service"
// 	"github.com/gin-gonic/gin"
// )

// var (
// 	videoService    service.VideoService       = service.New()
// 	videoController controller.VideoController = controller.New(videoService)
// )

// func main() {
// 	// initialize a new server
// 	server := gin.Default()

// 	server.GET("/posts", func(ctx *gin.Context) {
// 		ctx.JSON(200, videoController.FindAll())
// 	})
// 	server.POST("/posts", func(ctx *gin.Context) {
// 		ctx.JSON(200, videoController.Save(ctx))
// 	})

// 	server.Run(":8080")
// }
