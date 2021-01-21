package Schema

type OAuthRefreshToken struct {
	Id   			uint
	UserId 			int `gorm:"not null"`
	ClientId		int `gorm:"not null"`
	AccessTokenId 	int `gorm:"not null"`
	AccessToken 	string `gorm:"not null"`
	RefreshToken 	string `gorm:"not null"`
	Revoked 		bool `gorm:"not null"`
	CreatedAt 		string
	ExpiredAt 		string
}

func (b *OAuthRefreshToken) TableName() string {
	return "o_auth_refresh_tokens"
}
