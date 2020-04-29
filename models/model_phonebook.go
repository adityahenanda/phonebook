package models

import (
	"time"
)

//person entity
type Phonebook struct {
	PhonebookID int `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	FirstName   string
	LastName    string
	Address     []Address
	PhoneNumber []PhoneNumber
	CreatedAt   time.Time
	CreatedBy   string
	ModifiedAt  *time.Time
	ModifiedBy  string
	Deleted     bool
}

func (Phonebook) TableName() string {
	return "phonebooks"
}

//request data from frontend
type PhonebookRequest struct {
	PhonebookID int                  `json:"phonebookID"`
	FirstName   string               `json:"firstname" validate:"required"`
	LastName    string               `json:"lastname" validate:"required"`
	Address     []AddressRequest     `json:"address" validate:"required"`
	PhoneNumber []PhoneNumberRequest `json:"phoneNumber" validate:"required"`
	CreatedBy   string               `json:"createdBy"`
	ModifiedBy  string               `json:"modifiedBy"`
}

//request address
type AddressRequest struct {
	AddressID int    `json:"addressID"`
	Street    string `json:"street"`
	City      string `json:"city"`
	ZipCode   string `json:"zipCode"`
}

//phonenumber request
type PhoneNumberRequest struct {
	PhoneNumberID int    `json:"phoneNumberID"`
	PhoneNumber   string `json:"phoneNumber"`
}

//address entity
type Address struct {
	AddressID   int `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	PhonebookID int
	Street      string
	City        string
	ZipCode     string
	CreatedAt   time.Time
	CreatedBy   string
	ModifiedAt  *time.Time
	ModifiedBy  string
	Deleted     bool
}

//phonenumber entity
type PhoneNumber struct {
	PhoneNumberID int `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	PhonebookID   int
	PhoneNumber   string
	CreatedAt     time.Time
	CreatedBy     string
	ModifiedAt    *time.Time
	ModifiedBy    string
	Deleted       bool
}

type ResponsePhoneBooks struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Data      []Phonebook `json:"data"`
	TotalData int         `json:"totalData"`
}

type ResponsePhoneBookID struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	ID      int    `json:"id"`
}

type ResponsePhoneBook struct {
	Code    int       `json:"code"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    Phonebook `json:"data"`
}
