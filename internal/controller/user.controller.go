package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nk-hung/go-ecommerce-backend-api/internal/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	name := c.DefaultQuery("name", "isme")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "UserController 111... ping " + name, "uid": uid,
		"users": []string{
			"Hela", "Pexi", "Uza",
		},
		"user": uc.userService.GetUserInfo(),
	})
}
