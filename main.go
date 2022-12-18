package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	router.GET("/", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{"Message": "commit 4"})

	})

	router.Run(":7777")

}
