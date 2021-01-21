package Schema

import "time"

type User struct {
	Id 			uint
	FirstName 	string `gorm:"not null"`
	LastName 	string `gorm:"not null"`
	Email   	string `gorm:"unique;not null"`
	Password   	string `gorm:"not null"`
	Phone   	string `gorm:"not null"`
	Address 	*string
	CreatedAt   time.Time `sql:"DEFAULT:current_timestamp"`
	UpdatedAt   time.Time
}

func (b *User) TableName() string {
	return "users"
}
