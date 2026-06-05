package domain

import "time"

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}
