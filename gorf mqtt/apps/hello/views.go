package hello

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
	})
}
