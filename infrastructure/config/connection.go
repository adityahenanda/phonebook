package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func Init() *gorm.DB {

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.json")
	// viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	dbhost := viper.GetString(`database.host`)
	port := viper.GetString(`database.port`)
	username := viper.GetString(`database.user`)
	password := viper.GetString(`database.pass`)
	dbname := viper.GetString(`database.name`)
	// sslmode := viper.GetString(`database.ssl_mode`)

	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, dbhost, port, dbname)
	// dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s sslmode=%s", dbhost, port, username, dbname, password, sslmode)
	conn, err := gorm.Open("mysql", dbConnString)
	// defer conn.Close()
	if err != nil {
		fmt.Print(err)
	}

	Migrate(conn)

	return conn
}
