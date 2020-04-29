package usecase

import (
	"phonebook/models"
	"phonebook/service/repository"

	"github.com/jinzhu/gorm"
)

func GetPhonebooksUsecase(db *gorm.DB, limit int, page int) (data []models.Phonebook, total int, err error) {

	//transaction begin
	tx := db.Begin()

	//count all data
	totalData, err := repository.GetAllPhonebooks(tx)
	if err != nil {
		tx.Rollback()
		return data, total, err
	}
	total = len(totalData)

	//get single row phonebook
	res, err := repository.GetPhonebooks(tx, limit, page)
	if err != nil {
		tx.Rollback()
		return data, total, err
	}

	for _, item := range res {
		//get all phonenumber
		phonenumbers, err := repository.GetPhonenumberByPhoneID(tx, item.PhonebookID)
		if err != nil {
			tx.Rollback()
			return data, total, err
		}

		item.PhoneNumber = phonenumbers
		//get all addresses
		addresses, err := repository.GetAddressByPhoneID(tx, item.PhonebookID)
		if err != nil {
			tx.Rollback()
			return data, total, err
		}

		item.Address = addresses

		data = append(data, item)
	}

	tx.Commit()
	return data, total, err

}

func GetPhonebookUsecase(db *gorm.DB, phonebookID int) (data models.Phonebook, err error) {

	//transaction begin
	tx := db.Begin()

	//get single row phonebook
	data, err = repository.GetPhonebook(tx, phonebookID)
	if err != nil {
		tx.Rollback()
		return data, err
	}

	//get all phonenumber
	phonenumbers, err := repository.GetPhonenumberByPhoneID(tx, phonebookID)
	if err != nil {
		tx.Rollback()
		return data, err
	}

	//get all address
	addresses, err := repository.GetAddressByPhoneID(tx, phonebookID)
	if err != nil {
		tx.Rollback()
		return data, err
	}

	data.Address = addresses
	data.PhoneNumber = phonenumbers

	tx.Commit()
	return data, err

}

func CreatePhonebookUsecase(db *gorm.DB, req models.PhonebookRequest) (id int, err error) {

	//transaction begin
	tx := db.Begin()

	//create phonebook
	phonebookID, err := repository.CreatePhonebook(tx, &req)
	if err != nil {
		tx.Rollback()
		return phonebookID, err
	}

	//create address
	for _, address := range req.Address {
		_, err = repository.CreateAddress(tx, phonebookID, req.CreatedBy, &address)
		if err != nil {
			tx.Rollback()
			return phonebookID, err
		}
	}

	//create phonenumber
	for _, phonenumber := range req.PhoneNumber {
		_, err = repository.CreatePhonenumber(tx, phonebookID, req.CreatedBy, &phonenumber)
		if err != nil {
			tx.Rollback()
			return phonebookID, err
		}
	}

	tx.Commit()

	return phonebookID, err

}

//soft delete / set data to deleted flag = 1
func DeletePhonebookUsecase(db *gorm.DB, phonebookID int) (err error) {

	//transaction begin
	tx := db.Begin()

	//soft delete phone number by phonebookID
	err = repository.DeletePhonenumberByPhonebookID(tx, phonebookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	//soft delete address by phonebookID
	err = repository.DeleteAddressesByPhonebookID(tx, phonebookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	//soft delete phonebook
	err = repository.DeletePhonebook(tx, phonebookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

//delete data from db
func DestroyPhonebookUsecase(db *gorm.DB, phonebookID int) (err error) {

	//transaction begin
	tx := db.Begin()

	//soft delete phone number by phonebookID
	err = repository.DestroyPhonenumberByPhonebookID(tx, phonebookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	//soft delete address by phonebookID
	err = repository.DestroyAddressesByPhonebookID(tx, phonebookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	//soft delete phonebook
	err = repository.DestroyPhonebook(tx, phonebookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err

}

//soft delete then update data from db
func UpdatePhonebookUsecase(db *gorm.DB, req models.PhonebookRequest) (err error) {

	//transaction begin
	tx := db.Begin()

	//update all data to deleted = true (to handle missing phonenumber/address by request)
	//soft delete phone number by phonebookID
	err = repository.DeletePhonenumberByPhonebookID(tx, req.PhonebookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	//soft delete address by phonebookID
	err = repository.DeleteAddressesByPhonebookID(tx, req.PhonebookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	//soft delete phonebook
	err = repository.DeletePhonebook(tx, req.PhonebookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	//begin proses update all realated data
	for _, item := range req.PhoneNumber {
		//update phone number by phonebookID
		err = repository.UpdatePhonenumberByPhonebookID(tx, &item, req.PhonebookID, req.ModifiedBy)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, item := range req.Address {
		//updateaddress by phonebookID
		err = repository.UpdateAddressesByPhonebookID(tx, &item, req.PhonebookID, req.ModifiedBy)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//soft delete phonebook
	err = repository.UpdatePhonebook(tx, &req)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err

}
