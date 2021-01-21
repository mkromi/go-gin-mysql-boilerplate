package Schema

type OAuthAccessToken struct {
	Id   		uint
	UserId 		int `gorm:"not null"`
	ClientId	int `gorm:"not null"`
	AccessToken string `gorm:"not null"`
	Revoked 	bool `gorm:"not null"`
	CreatedAt 	string
	ExpiredAt 	string
}

func (b *OAuthAccessToken) TableName() string {
	return "o_auth_access_tokens"
}
