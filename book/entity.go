package book

import "time"

type Book struct {
	ID          int    `gorm:"type:int;AUTO_INCREMENT;PRIMARY_KEY"`
	Title       string `gorm:"type:varchar(100);unique;not null"`
	Description string `gorm:"type:text"`
	Price       int    `gorm:"type:int;not null"`
	Rating      int    `gorm:"type:int;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
