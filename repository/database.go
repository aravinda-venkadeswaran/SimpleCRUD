package repository

import (
	"errors"
	"localSandbox/SimpleCRUD/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

func InitializeDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Address{}, &models.Company{}, &models.Geo{}, &models.Post{}, &models.Comment{})
	DB = db
	return db
}

func SaveToDB(data interface{}) {
	DB.Save(data)
}

func FindAll(params ...interface{}) {
	data := params[0]
	if len(params) == 2 {
		preload := params[1].(string)
		DB.Preload(preload).Preload(clause.Associations).Find(data)
	} else {
		DB.Preload(clause.Associations).Find(data)
	}
}

func FindByID(params ...interface{}) (interface{}, error) {
	id := params[1].(int)
	data := params[0]
	if len(params) == 3 {
		preload := params[2].(string)
		if err := DB.Where("id = ?", id).Preload(preload).Preload(clause.Associations).First(data).Error; err != nil {
			return data, errors.New("Record not found")
		}
	} else {
		if err := DB.Where("id = ?", id).Preload(clause.Associations).First(data).Error; err != nil {
			return data, errors.New("Record not found")
		}
	}
	return data, nil
}

func DeleteFromDB(data interface{}, id int) {
	DB.Delete(data, id)
}
