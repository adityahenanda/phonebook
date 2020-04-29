package repository

import (
	"phonebook/models"
	"time"

	"github.com/jinzhu/gorm"
)

func CreatePhonebook(db *gorm.DB, phonebookRequest *models.PhonebookRequest) (id int, err error) {

	row := new(models.Phonebook)
	var phonebook models.Phonebook
	phonebook.CreatedAt = time.Now()
	phonebook.FirstName = phonebookRequest.FirstName
	phonebook.LastName = phonebookRequest.LastName
	phonebook.CreatedBy = phonebookRequest.CreatedBy
	err = db.Create(&phonebook).Scan(row).Error
	if err != nil {
		return row.PhonebookID, err
	}

	return row.PhonebookID, err
}

func GetPhonebook(db *gorm.DB, phonebookID int) (data models.Phonebook, err error) {

	rows, err := db.Raw(`SELECT * FROM phonebooks where phonebook_id = ?`, phonebookID).Rows()
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		err = db.ScanRows(rows, &data)
		if err != nil {
			return data, err
		}

	}
	return data, err
}

func GetPhonebooks(db *gorm.DB, limit int, page int) (data []models.Phonebook, err error) {

	//default limit, offset
	offset := 0
	if limit == 0 {
		limit = 100
	}
	if page > 0 {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw(`SELECT * FROM phonebooks order by phonebook_id asc LIMIT ? OFFSET ?`, limit, offset).Rows()
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.Phonebook
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return data, err
		}

		data = append(data, temp)

	}
	return data, err
}

func GetAllPhonebooks(db *gorm.DB) (data []models.Phonebook, err error) {

	rows, err := db.Raw(`SELECT * FROM phonebooks`).Rows()
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.Phonebook
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return data, err
		}

		data = append(data, temp)

	}
	return data, err
}

//soft delete
func DeletePhonebook(db *gorm.DB, phonebookID int) (err error) {

	err = db.Exec(`Update phonebooks set deleted = 1 where phonebook_id = ?`, phonebookID).Error
	if err != nil {
		return err
	}

	return err
}

//delete data from db
func DestroyPhonebook(db *gorm.DB, phonebookID int) (err error) {

	err = db.Exec(`delete from phonebooks where phonebook_id = ?`, phonebookID).Error
	if err != nil {
		return err
	}

	return err
}

//update phonebook
func UpdatePhonebook(db *gorm.DB, phonebookRequest *models.PhonebookRequest) (err error) {

	err = db.Exec(`update phonebooks 
	set first_name = ?, last_name = ? ,modified_at = ?, modified_by =? , deleted = 0
	where phonebook_id = ?`, phonebookRequest.FirstName, phonebookRequest.LastName, time.Now(), phonebookRequest.ModifiedBy, phonebookRequest.PhonebookID).Error
	if err != nil {
		return err
	}

	return err
}
