package api

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	v1 := r.Group("/user")
	v1.POST("/register")

	return r

}
