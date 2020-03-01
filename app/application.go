package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApplication is a function to be called from main.go for starting the apllication
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
