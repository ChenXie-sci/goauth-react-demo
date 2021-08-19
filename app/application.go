package app

import (
	"github.com/gin-genic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Run(":8000")
}
