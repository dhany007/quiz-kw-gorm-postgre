package models

import (
	"fmt"
	"time"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) Print() {
	fmt.Println("ID :", p.ID)
	fmt.Println("Name :", p.Name)
	fmt.Println("Brand :", p.Brand)
	fmt.Println("UserID :", p.UserID)
}
