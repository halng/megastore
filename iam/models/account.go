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
	CreateAt  int64     `json:"createAt"`
	UpdateAt  int64     `json:"updateAt"`
	CreateBy  string    `json:"createBy"`
	UpdateBy  string    `json:"updateBy"`
}

func (account *Account) SaveAccount() (*Account, error) {

	account.ID = uuid.New()
	account.CreateAt = time.Now().Unix()
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
	account.UpdateAt = time.Now().Unix()
	account.UpdateBy = account.Username
	return nil
}

func ExistsByEmailOrUsername(email string, username string) (Account, error) {
	var account Account
	err := DB.Model(&Account{}).Where("email = ? OR username = ?", email, username).Take(&account).Error
	return account, err
}

func (account *Account) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	return err == nil
}

func (account *Account) GenerateAccessToken() string {
	token, err = token.Ge
}
