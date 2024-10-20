package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tanhaok/megastore/constants"
	"github.com/tanhaok/megastore/db"
	"github.com/tanhaok/megastore/dto"
	"github.com/tanhaok/megastore/logging"
	"github.com/tanhaok/megastore/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	RoleId    string    `json:"roleId"`
	CreateAt  int64     `json:"createAt"`
	UpdateAt  int64     `json:"updateAt"`
	CreateBy  string    `json:"createBy"`
	UpdateBy  string    `json:"updateBy"`
}

func (account *Account) SaveAccount() (*Account, error) {

	account.ID = uuid.New()
	account.CreateAt = time.Now().Unix()
	if account.CreateBy != "" {
		account.CreateBy = constants.DefaultCreator
	}

	if err := db.DB.Postgres.Create(&account).Error; err != nil {
		return &Account{}, err
	}

	return account, nil

}

func (account *Account) BeforeSave() error {
	account.UpdateAt = time.Now().Unix()
	account.UpdateBy = account.Username
	return nil
}

func ExistsByEmailOrUsername(email string, username string) bool {
	var count int64
	db.DB.Postgres.Model(&Account{}).Where("email = ? OR username = ?", email, username).Count(&count)
	return count > 0
}

func GetAccountByUsername(username string) (Account, error) {
	var account Account
	err := db.DB.Postgres.Model(&Account{}).Where(" username = ?", username).Take(&account).Error
	return account, err
}

func (account *Account) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	return err == nil
}

func (account *Account) GenerateAccessToken() string {

	role, err := GetRoleById(account.RoleId)
	if err != nil {
		logging.LOGGER.Error("Cannot get role for user %s ", zap.Any("account", account))
		return ""
	}
	jwtToken, err := utils.GenerateJWT(account.ID.String(), account.Username, role)
	if err != nil {
		log.Printf("Cannot generate access token for user %s ", account.Username)
		return ""
	}

	hashedMD := utils.ComputeMD5([]string{account.ID.String()})
	cacheId := uuid.New().String()
	err = db.SaveDataToCache(fmt.Sprintf("%s_%s", hashedMD, cacheId), jwtToken)

	if err != nil {
		log.Printf("Cannot save token in cache")
		return ""
	}

	return cacheId
}

func (account *Account) GetSerializedMessageForActiveNewUser() string {
	var activeNewUser dto.ActiveNewUser
	activeNewUser.Username = account.Username
	activeNewUser.Email = account.Email
	activeNewUser.Token = utils.ComputeHMAC256(account.Username, account.Email)
	activeNewUser.ExpiredTime = time.Now().UnixMilli() + 1000*60*60*24 // 1 day

	serialized, err := json.Marshal(activeNewUser)
	if err != nil {
		log.Printf("Cannot serialize data")
		return ""
	}
	return string(serialized)
}
