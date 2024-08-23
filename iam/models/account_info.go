package models

// AccountInfo struct
type AccountInfo struct {
	ID               string `json:"id"`
	AccountId        string `json:"account_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	PhoneNumber      string `json:"phone_number"`
	PrimaryAddress   string `json:"primary_address"`
	SecondaryAddress string `json:"secondary_address"`
}
