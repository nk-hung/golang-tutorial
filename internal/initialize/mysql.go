package initialize

import (
	"fmt"
	"time"

	"github.com/nk-hung/go-ecommerce-backend-api/global"
	"github.com/nk-hung/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CheckErrorPanic(err error, msg string) {
	if err != nil {
		global.Logger.Error(msg, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql
	global.Logger.Info("MySQL config", zap.Any("config", m))
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	CheckErrorPanic(err, "Failed to connect to MySQL database ::: "+dsn)
	// Set the connection pool

	global.Logger.Info("Connected to MySQL database successfully")
	global.Mdb = db

	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	CheckErrorPanic(err, "Failed to get database connection pool")
	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
	// sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(&po.User{}, &po.Role{})
	CheckErrorPanic(err, "Failed to migrate tables")
	// global.Logger.Info("Tables migrated successfully")
	// global.Mdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&po.User{}, &po.Role{})
	global.Logger.Info("Tables migrated successfully")
}
