package Services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go-gin-mysql-boilerplate/Config"
	"go-gin-mysql-boilerplate/Models"
	"go-gin-mysql-boilerplate/Models/Schema"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type accessToken struct {
	Token interface{}
	CreatedAt interface{}
	ExpiredAt interface{}
}

type refreshToken struct {
	Token interface{}
	CreatedAt interface{}
	ExpiredAt interface{}
}

type TokenResponse struct {
	AccessToken accessToken
	RefreshToken refreshToken
}


func GenerateTokens(userId string, clientName string) TokenResponse {
	tokenForAccess, errAccess := bcrypt.GenerateFromPassword([]byte("accessToken" + time.Now().String() + userId), bcrypt.DefaultCost)
	if errAccess != nil {
		panic(errAccess)
	}

	tokenForRefresh, errRefresh := bcrypt.GenerateFromPassword([]byte("refreshToken" + time.Now().String() + userId), bcrypt.DefaultCost)
	if errRefresh != nil {
		panic(errRefresh)
	}

	tokens := TokenResponse{
		AccessToken: accessToken{
			Token: string(tokenForAccess),
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			ExpiredAt: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
		},
		RefreshToken: refreshToken{
			Token:     string(tokenForRefresh),
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			ExpiredAt: time.Now().Add(time.Hour * 72).Format("2006-01-02 15:04:05"),
		},
	}

	var clientUser Schema.OAuthClient
	if err := Config.DB.Raw("select * from o_auth_clients " +
		"where name = ?", clientName).Scan(&clientUser).Error; err != nil {
		return TokenResponse{}
	}

	intUserId, _ := strconv.Atoi(userId)

	var saveAccessToken Schema.OAuthAccessToken
	saveAccessToken.UserId 		= intUserId
	saveAccessToken.ClientId	= int(clientUser.Id)
	saveAccessToken.AccessToken	= fmt.Sprintf("%v", tokens.AccessToken.Token)
	saveAccessToken.Revoked		= false
	saveAccessToken.CreatedAt	= fmt.Sprintf("%v", tokens.AccessToken.CreatedAt)
	saveAccessToken.ExpiredAt	= fmt.Sprintf("%v", tokens.AccessToken.ExpiredAt)

	saveAccessTokenErr := Models.AuthAccessTokenCreate(&saveAccessToken)
	if saveAccessTokenErr != nil {
		return TokenResponse{}
	}

	var saveRefreshToken Schema.OAuthRefreshToken
	saveRefreshToken.UserId			= intUserId
	saveRefreshToken.ClientId		= int(clientUser.Id)
	saveRefreshToken.AccessTokenId 	= int(saveAccessToken.Id)
	saveRefreshToken.AccessToken 	= saveAccessToken.AccessToken
	saveRefreshToken.RefreshToken 	= fmt.Sprintf("%v", tokens.RefreshToken.Token)
	saveRefreshToken.Revoked		= false
	saveRefreshToken.CreatedAt		= fmt.Sprintf("%v", tokens.RefreshToken.CreatedAt)
	saveRefreshToken.ExpiredAt		= fmt.Sprintf("%v", tokens.RefreshToken.ExpiredAt)

	saveRefreshTokenErr := Models.AuthRefreshTokenCreate(&saveRefreshToken)
	if saveRefreshTokenErr != nil {
		return TokenResponse{}
	}

	return tokens
}

func Signout(token string) bool {
	if err := Config.DB.
		Table("o_auth_refresh_tokens").
		Where("access_token = ?", token).
		Update("revoked", true).
		Error; err != nil {
		return false
	}

	if err := Config.DB.
		Table("o_auth_access_tokens").
		Where("access_token = ?", token).
		Update("revoked", true).
		Error; err != nil {
		return false
	}
	return true
}

func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}