package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"

	"phonebook/infrastructure/config"
	_httpDeliver "phonebook/service/delivery/http_handler"
)

func main() {
	db := config.Init()

	route := gin.Default()

	_httpDeliver.NewHttpHandler(route, db)

	route.Run(viper.GetString("server.address"))
}
