package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/yushuailiu/easygin/pkg/app"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/go-ini/ini"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID		uint		`gorm:"primary_key"`
	CreatedAt	time.Time	`sql:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	UpdatedAt	time.Time	`sql:"default:CURRENT_TIMESTAMP"`
}

func Bootstrap() {
	mysqlConfig, err := app.GetConfig().GetSection("mysql")
	if err != nil {
		panic(err)
	}

	initMysql(mysqlConfig)
}

func initMysql(config *ini.Section) {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Key("username").String(),
		config.Key("password").String(),
		config.Key("host").String(), config.Key("dbname"))
	gormDB, err := gorm.Open("mysql", url)

	maxIdle, err := config.Key("maxIdle").Int()

	if err != nil {
		panic(err)
	}

	gormDB.DB().SetMaxIdleConns(maxIdle)
	if err != nil {
		panic(err)
	}
	DB = gormDB
}
