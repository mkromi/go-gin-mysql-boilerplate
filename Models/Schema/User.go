package Schema

import "time"

type User struct {
	Id 			uint `json:"Id"`
	FirstName 	string `json:"FirstName" gorm:"not null"`
	LastName 	string `json:"LastName" gorm:"not null"`
	Email   	string `json:"Email" gorm:"unique;not null"`
	Password   	string `json:"Password" gorm:"not null"`
	Phone   	string `json:"Phone" gorm:"not null"`
	Address 	*string `json:"Address"`
	CreatedAt   time.Time `sql:"DEFAULT:current_timestamp"`
	UpdatedAt   time.Time
}

func (b *User) TableName() string {
	return "users"
}
