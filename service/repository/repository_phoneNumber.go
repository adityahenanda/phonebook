package repository

import (
	"phonebook/models"
	"time"

	"github.com/jinzhu/gorm"
)

func CreatePhonenumber(db *gorm.DB, phoneID int, createdBy string, req *models.PhoneNumberRequest) (id int, err error) {

	row := new(models.PhoneNumber)
	var phoneNumber models.PhoneNumber
	phoneNumber.CreatedAt = time.Now()
	phoneNumber.PhoneNumber = req.PhoneNumber
	phoneNumber.PhonebookID = phoneID
	phoneNumber.CreatedBy = createdBy
	err = db.Create(&phoneNumber).Scan(row).Error
	if err != nil {
		return phoneNumber.PhoneNumberID, err
	}

	return phoneNumber.PhoneNumberID, err
}

func GetPhonenumberByPhoneID(db *gorm.DB, phonebookID int) (data []models.PhoneNumber, err error) {
	rows, err := db.Raw(`SELECT * FROM phone_numbers where phonebook_id = ?`, phonebookID).Rows()
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.PhoneNumber
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	return data, err
}

//soft delete
func DeletePhonenumberByPhonebookID(db *gorm.DB, phonebookID int) (err error) {

	err = db.Exec(`Update phone_numbers set deleted = 1 where phonebook_id = ?`, phonebookID).Error
	if err != nil {
		return err
	}

	return err
}

//delete data from db
func DestroyPhonenumberByPhonebookID(db *gorm.DB, phonebookID int) (err error) {

	err = db.Exec(`delete from phone_numbers where phonebook_id = ?`, phonebookID).Error
	if err != nil {
		return err
	}

	return err
}

func UpdatePhonenumberByPhonebookID(db *gorm.DB, req *models.PhoneNumberRequest, phonebookID int, modifiedBy string) (err error) {

	err = db.Exec(`update phone_numbers 
	set phone_number = ? , modified_at = ?, modified_by =? , deleted = 0
	where phonebook_id = ? and phone_number_id = ?`, req.PhoneNumber, time.Now(), modifiedBy, phonebookID, req.PhoneNumberID).Error
	if err != nil {
		return err
	}

	return err
}
