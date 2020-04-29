package config

import (
	. "phonebook/models"

	"github.com/jinzhu/gorm"
)

func Migrate(idb *gorm.DB) {

	idb.Debug().AutoMigrate(
		&Phonebook{},
		&Address{},
		&PhoneNumber{},
	)

}
