package http_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type InDB struct {
	DB *gorm.DB
}

func NewHttpHandler(route *gin.Engine, db *gorm.DB) {

	handler := &InDB{DB: db}

	v1 := route.Group("/v1")
	{
		api := v1.Group("/api")
		{
			phonebook := api.Group("/phonebook")
			{
				phonebook.GET("", handler.GetPhonebooks)
				phonebook.GET("/:phoneBookID", handler.GetPhonebook)
				phonebook.PUT("/:phoneBookID", handler.UpdatePhonebook)
				phonebook.POST("", handler.CreatePhonebook)
				phonebook.DELETE("/:phoneBookID", handler.DeletePhonebook)
				phonebook.DELETE("/:phoneBookID/destroy", handler.DestroyPhonebook)
			}
		}
	}
}
