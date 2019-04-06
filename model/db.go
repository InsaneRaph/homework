package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"homeworkprojet/config"
	"homeworkprojet/utils"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	var err error
	DB, err = gorm.Open("mysql", config.AppConfig.DataSource)
	utils.PanicOnError(errors.Wrap(err, "failed to connect to DB"))
	DB.AutoMigrate(&User{}, &CardScheme{})

	return DB
}
