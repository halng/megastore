package models

import (
	"github.com/google/uuid"
	"github.com/tanhaok/megastore/db"
)

const (
	RoleAnonymous  = "anonymous"
	RoleAdmin      = "admin"
	RoleStaff      = "staff"
	RoleUser       = "user"
	RoleSuperAdmin = "super_admin"
)

type Role struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Account []Account
}

func GetRoleIdByName(name string) (string, error) {
	var roleID string
	err := db.DB.Postgres.Model(&Role{}).Where("name = ?", name).Select("id").Row().Scan(&roleID)
	if err != nil {
		return "", err
	}
	return roleID, nil
}

func GetRoleById(id string) (string, error) {
	var role Role
	err := db.DB.Postgres.Where("id = ?", id).First(&role).Error
	if err != nil {
		return "", err
	}
	return role.Name, nil
}
