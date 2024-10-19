package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tanhaok/megastore/db"
	"github.com/tanhaok/megastore/logging"
	"go.uber.org/zap"
	"os"
)

func Initialize() {
	DB := db.DB
	DB.Postgres.AutoMigrate(&Account{})
	DB.Postgres.AutoMigrate(&Role{})
	initRole(DB.Postgres)
	initMasterUser()

}

// InitRole auto create roles when the app starts
func initRole(db *gorm.DB) {
	roles := []string{RoleAnonymous, RoleAdmin, RoleStaff, RoleUser, RoleSuperAdmin}

	for _, roleName := range roles {
		role := Role{ID: uuid.New(), Name: roleName}
		if err := db.Create(&role).Error; err != nil {
			logging.LOGGER.Error("Error occurred when creating role",
				zap.String("roleName", roleName),
				zap.Any("err", err))

		} else {
			logging.LOGGER.Info("Role created successfully", zap.String("roleName", roleName))
		}
	}

}

func initMasterUser() {
	masterUsername := os.Getenv("MASTER_USERNAME")
	masterPassword := os.Getenv("MASTER_PASSWORD")
	masterEmail := os.Getenv("MASTER_EMAIL")
	masterFirstName := os.Getenv("MASTER_FIRST_NAME")
	masterLastName := os.Getenv("MASTER_LAST_NAME")

	masterAccount := Account{Username: masterUsername, Password: masterPassword, Email: masterEmail, FirstName: masterFirstName, LastName: masterLastName}
	_, err := masterAccount.SaveAccount()
	if err != nil {
		logging.LOGGER.Error("Cannot create master account", zap.Any("error", err))
		panic("Cannot create master account")
	}

	logging.LOGGER.Info("Master account created successfully")
}
