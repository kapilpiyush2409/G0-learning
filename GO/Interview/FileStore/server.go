package main

import (
	"filestore/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/store/files", controllers.AddFile)
	router.GET("/store/files", controllers.ListFiles)
	router.DELETE("/store/files/:filename", controllers.RemoveFile)
	router.PUT("/store/files/:filename", controllers.UpdateFile)
	router.POST("/store/upload", controllers.UploadMultipleTextFiles)
	router.GET("/store/wc", controllers.CountWordsInDirectory)
	router.GET("/store/freqwords", controllers.WordsFrequencyInDirectory)
	gin.SetMode(gin.DebugMode)
	router.Run(":8080")
}
