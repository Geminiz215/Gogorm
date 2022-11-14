package users

import "time"

type UserRequestCreate struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRequestFindOne struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username" `
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
