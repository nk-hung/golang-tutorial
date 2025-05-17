package initialize

import (
	"fmt"

	"github.com/nk-hung/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Global config:::", global.Config.Mysql.DbName)
	InitLogger()
	global.Logger.Info("Logger initialized", zap.String("test", "logger ok"))
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8002")
}
