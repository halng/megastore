package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tanhaok/MyStore/db"
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
			fmt.Printf("Error creating role %s: %v", roleName, err)

		} else {
			fmt.Printf("Role %s created successfully", roleName)
		}
	}

}
