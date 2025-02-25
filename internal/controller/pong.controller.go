package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "isme")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong controller ... 1111 " + name, "uid": uid,
		"users": []string{
			"Hela", "Pexi", "Uza",
		},
	})
}
