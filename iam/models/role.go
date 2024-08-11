package models

import (
	"github.com/google/uuid"
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
