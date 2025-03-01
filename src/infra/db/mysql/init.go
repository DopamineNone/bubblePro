package mysql

import (
	"fmt"
	"github.com/DopamineNone/bubblePro/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error

	conf := config.GetConf().MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Database)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(dsn)
	if err != nil {
		panic("failed to connect database")
	}
}
