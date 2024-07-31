package models

import (
	"github.com/google/uuid"
	"github.com/tanhaok/MyStore/constants"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// TODO: add relationship with role

type Account struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreateAt  string    `json:"createAt"`
	UpdateAt  string    `json:"updateAt"`
	CreateBy  string    `json:"createBy"`
	UpdateBy  string    `json:"updateBy"`
}

func (account *Account) SaveAccount() (*Account, error) {

	account.ID = uuid.New()
	account.CreateAt = time.Now().String()
	account.CreateBy = constants.DefaultCreator

	if err := DB.Create(&account).Error; err != nil {
		return &Account{}, err
	}

	return account, nil

}

func (account *Account) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	account.Password = string(hashedPassword)
	account.UpdateAt = time.Now().String()
	account.UpdateBy = account.Username
	return nil
}
