package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"primary_key" gorm:"index" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string
	Password string
	Role     Role
	Email    string
	CreateAt string
	UpdateAt string
	CreateBy string
	UpdateBy string
}

//
//func ExistAccountByID(id string) bool {
//	var account Account
//	err := db.Select("id").Where("id = ?", id).First(&account).Error()
//	if account.ID != uuid.Nil {
//		return true
//	}
//
//	return false
//}
