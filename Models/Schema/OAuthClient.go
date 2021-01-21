package Schema

import "time"

type OAuthClient struct {
	Id 			uint
	Name 		string `gorm:"not null"`
	Secret 		string `gorm:"not null"`
	CreatedAt   time.Time `sql:"DEFAULT:current_timestamp"`
}

func (b *OAuthClient) TableName() string {
	return "o_auth_clients"
}
