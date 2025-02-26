package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nk-hung/go-ecommerce-backend-api/internal/services"
	"github.com/nk-hung/go-ecommerce-backend-api/pkg/response"
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
	// c.JSON(http.StatusOK, response.ResponseData{
	// 	Code:    20001,
	// 	Message: "Thanh cong",
	// 	Data:    "sdaf",
	// })
	response.SuccessReponse(c, 20001, []string{"Paki", "ct4"})
}
