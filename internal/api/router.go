package api

import (
	"cloud_storage/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(u *handlers.UserHandler) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	v1 := r.Group("/user")
	v1.POST("/register", u.Register)
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello User",
		})
	})

	return r

}
