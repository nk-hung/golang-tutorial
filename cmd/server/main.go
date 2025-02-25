package main

import (
	"github.com/nk-hung/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRoute()

	r.Run(":8998")
}
