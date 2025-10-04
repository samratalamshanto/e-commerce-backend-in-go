package user

import "time"

type CreateUserDTO struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"email"`
	Phone     string `json:"phone"`
	AccountID string `json:"accountId" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

type UpdateUserDTO struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"email"`
	Phone     string `json:"phone"`
	AccountID string `json:"accountId" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

type LoginDTO struct {
	AccountID string `json:"accountId" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	AccountID string    `json:"accountId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
