package transfer

import "github.com/gin-gonic/gin"

var router *gin.Engine

func NewRouter() {
	router = gin.Default()

	router.GET("/text", func(c *gin.Context) {

	})

	router.Run()
}