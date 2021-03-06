package main

import (
	"github.com/spf13/pflag"

	"github.com/yushuailiu/gracefulServer"
	"github.com/yushuailiu/easygin/pkg/app"
	"github.com/yushuailiu/easygin/pkg/logging"
	"github.com/yushuailiu/easygin/app/models"
	"github.com/yushuailiu/easygin/pkg/database/mysql"
	"github.com/yushuailiu/easygin/pkg/database/redis"
)

func main() {
	var env *string = pflag.String("env", "", "运行环境类型：development、sandbox、production")
	pflag.Parse()
	if *env != "sandbox" && *env != "production" && *env != "development" {
		pflag.PrintDefaults()
		return
	}

	routes := app.Bootstrap(*env)

	// start mysql
	mysql.Bootstrap()

	// start redis
	redis.Bootstrap()

	result := redis.RedisClient.Get("liu")
	println(result.Val())

	models.AddUser("shuai", "liu")

	// start logging
	logging.Bootstrap()

	graceful := gracefulServer.NewGraceful()
	err := graceful.ListenAndServer(":8082", routes)
	if err != nil {
		panic(err)
	}
}
