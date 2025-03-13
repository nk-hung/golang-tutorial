package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	c "github.com/nk-hung/go-ecommerce-backend-api/internal/controller"
	"github.com/nk-hung/go-ecommerce-backend-api/internal/middlewares"
)

func AAA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ->> AAA")
		c.Next()
		fmt.Println("After ->> AAA")
	}
}

func BBB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ->> BBB")
		c.Next()
		fmt.Println("After ->> BBB")
	}
}

func CCC() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ->> CCC")
		c.Next()
		fmt.Println("After ->> CCC")
	}
}

func NewRoute() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware(), AAA(), BBB(), CCC())

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping/:name", c.NewPongController().Pong)
		v1.GET("/pong", c.NewUserController().GetUserById)
		v1.POST("/user/1", c.NewUserController().GetUserById)
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
