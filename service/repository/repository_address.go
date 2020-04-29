package repository

import (
	"phonebook/models"
	"time"

	"github.com/jinzhu/gorm"
)

func CreateAddress(db *gorm.DB, phoneID int, createdBy string, req *models.AddressRequest) (id int, err error) {

	row := new(models.Address)
	var address models.Address
	address.CreatedAt = time.Now()
	address.CreatedBy = createdBy
	address.City = req.City
	address.Street = req.Street
	address.ZipCode = req.ZipCode
	address.PhonebookID = phoneID
	err = db.Create(&address).Scan(row).Error
	if err != nil {
		return address.AddressID, err
	}

	return address.AddressID, err

}

func GetAddressByPhoneID(db *gorm.DB, phonebookID int) (data []models.Address, err error) {
	rows, err := db.Raw(`SELECT * FROM addresses where phonebook_id = ?`, phonebookID).Rows()
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.Address
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return data, err
		}

		data = append(data, temp)

	}

	return data, err
}

//soft delete
func DeleteAddressesByPhonebookID(db *gorm.DB, phonebookID int) (err error) {

	err = db.Exec(`Update addresses set deleted = 1 where phonebook_id = ?`, phonebookID).Error
	if err != nil {
		return err
	}

	return err
}

//delete data from db
func DestroyAddressesByPhonebookID(db *gorm.DB, phonebookID int) (err error) {

	err = db.Exec(`delete from addresses where phonebook_id = ?`, phonebookID).Error
	if err != nil {
		return err
	}

	return err
}

func UpdateAddressesByPhonebookID(db *gorm.DB, req *models.AddressRequest, phonebookID int, modifiedBy string) (err error) {

	err = db.Exec(`update addresses 
	set street = ?, city = ?, zip_code = ? , modified_at = ? , modified_by = ? , deleted = 0
	where phonebook_id = ? and address_id = ?`, req.Street, req.City, req.ZipCode, time.Now(), modifiedBy, phonebookID, req.AddressID).Error
	if err != nil {
		return err
	}

	return err
}
