package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

const (
	RoleAnonymous  = "anonymous"
	RoleAdmin      = "admin"
	RoleStaff      = "staff"
	RoleUser       = "user"
	RoleSuperAdmin = "super_admin"
)

type Role struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// InitRole auto create roles when the app starts
func initRole(db *gorm.DB) {
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
