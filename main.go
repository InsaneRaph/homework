package main

import (
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"homeworkprojet/config"
	"homeworkprojet/model"
	"homeworkprojet/router"
	"homeworkprojet/utils"
	"net/http"
	"strconv"
)

func main() {

	level, err := log.ParseLevel(config.AppConfig.LogLevel)
	utils.PanicOnError(err)
	log.SetLevel(level)

	db := model.InitDb()
	defer db.Close()

	listen := ":" + strconv.Itoa(config.AppConfig.Port)
	log.Infoln("listening on ", listen)

	err = http.ListenAndServe(listen, cors.Default().Handler(router.RegisterRoutes()))

	utils.PanicOnError(err)

}
