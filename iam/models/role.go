package models

const (
	RoleAnonymous  = "anonymous"
	RoleAdmin      = "admin"
	RoleStaff      = "staff"
	RoleUser       = "user"
	RoleSuperAdmin = "super_admin"
)

type Role struct {
	ID   string
	Name string
}
