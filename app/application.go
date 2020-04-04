package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shchaslyvyi/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

// StartApplication is a function to be called from main.go for starting the apllication
func StartApplication() {
	mapUrls()
	logger.Info("About to start the application...")
	router.Run(":8081")
}
