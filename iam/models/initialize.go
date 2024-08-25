package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tanhaok/megastore/db"
	"github.com/tanhaok/megastore/logging"
	"go.uber.org/zap"
)

func Initialize() {
	DB := db.DB
	DB.Postgres.AutoMigrate(&Account{})
	DB.Postgres.AutoMigrate(&Role{})
	InitRole(DB.Postgres)
}

// InitRole auto create roles when the app starts
func InitRole(db *gorm.DB) {
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
