package model

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
	MongoDB     *mongo.Client
)

func InitDB() {
	// 构建连接
	dsn := strings.Join([]string{
		viper.GetString("mysql.user"), ":", viper.GetString("mysql.password"),
		"@tcp(", viper.GetString("mysql.ip"),
		":",
		viper.GetString("mysql.port"),
		")/",
		viper.GetString("mysql.database"),
		"?charset=",
		viper.GetString("mysql.charset"),
		"&parseTime=",
		viper.GetString("mysql.parseTime"),
		"&loc=",
		viper.GetString("mysql.loc"),
	}, "")

	// 打开数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 错误处理
	if err != nil {
		fmt.Println("=> 数据库初始化失败...")
		fmt.Println("=> ", err)
		panic(err)
	}

	fmt.Println("=> 数据库初始化成功...")
	DB = db
	migration()
}

func migration() {
	CreateDbUserBasic()
}
