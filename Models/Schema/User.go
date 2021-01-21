package Schema

import "time"

type User struct {
	Id 			uint `json:"Id"`
	FirstName 	string `json:"FirstName" gorm:"not null" binding:"required"`
	LastName 	string `json:"LastName" gorm:"not null" binding:"required"`
	Email   	string `json:"Email" gorm:"unique;not null" binding:"required"`
	Password   	string `json:"Password" gorm:"not null" binding:"required"`
	Phone   	string `json:"Phone" gorm:"not null" binding:"required"`
	Address 	*string `json:"Address"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b *User) TableName() string {
	return "users"
}
