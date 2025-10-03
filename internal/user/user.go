package user

import "ecom-backend/internal/common"

type User struct {
	common.BaseEntity
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AccountId string `json:"accountId"`
	Password  string `json:"-"`
}
