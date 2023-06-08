package hello

import "github.com/gin-gonic/gin"

func Urls(r *gin.Engine) {
	r.GET("/hello", Hello)
}
