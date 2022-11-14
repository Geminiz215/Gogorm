package users

import "time"

type User struct {
	ID        uint   `gorm:"type:int;AUTO_INCREMENT;PRIMARY_KEY"`
	Username  string `gorm:"type:varchar(50);unique;not null"`
	Email     string `gorm:"type:varchar(50);not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
