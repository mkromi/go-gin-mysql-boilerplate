package Models

import (
	"go-gin-mysql-boilerplate/Config"
	"go-gin-mysql-boilerplate/Models/Schema"
)

func AuthAccessTokenCreate(data *Schema.OAuthAccessToken) (err error) {
	if err = Config.DB.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func AuthRefreshTokenCreate(data *Schema.OAuthRefreshToken) (err error) {
	if err = Config.DB.Create(data).Error; err != nil {
		return err
	}
	return nil
}
