package api

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	//v1 := r.Group("/api/v1")

	return r
}
