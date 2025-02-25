package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping/:name", Pong)
		v1.GET("/pong", Pong)
		v1.POST("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
		v1.DELETE("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "isme")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong ... ping " + name, "uid": uid,
		"users": []string{
			"Hela", "Pexi", "Uza",
		},
	})
}
