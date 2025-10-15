package handlers

import (
	"cloud_storage/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service *services.User
}

func NewUserHandler(service *services.User) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) Register(c *gin.Context) {
	var req UserReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect JSON"})
		return
	}

	if err := u.service.Register(req.Username, req.Password, req.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is already taken"})
		return

	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})

}

func (u *UserHandler) Login(c *gin.Context) {
	var req UserReq

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect JSON"})
	}

	if err := u.service.Login(req.Username, req.Password); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "User!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in!"})

}

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
