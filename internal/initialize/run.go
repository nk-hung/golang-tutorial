package initialize

import (
	"fmt"

	"github.com/nk-hung/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Global config:::", global.Config.Mysql.DbName)
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8002")
}
