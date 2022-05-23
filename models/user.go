package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Product   []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) Print() {
	fmt.Println("ID :", user.ID)
	fmt.Println("Email :", user.Email)
}
