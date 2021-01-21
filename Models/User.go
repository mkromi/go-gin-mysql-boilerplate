package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"go-gin-mysql-boilerplate/Config"
	"go-gin-mysql-boilerplate/Models/Schema"
	"go-gin-mysql-boilerplate/Validations"
)

func UserFetchAll(user *[]Schema.User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func UserFetchSingle(user *Schema.User, userId string) (err error) {
	if err = Config.DB.Where("id = ?", userId).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UserCreate(request *Validations.UserCreate) (err error) {
	if err = Config.DB.Table("users").Create(request).Error; err != nil {
		return err
	}
	return nil
}

func UserUpdate(request *Validations.UserUpdate, userId string) (err error) {
	if err = Config.DB.Table("users").Where("id = ?", userId).Update(request).Error; err != nil {
		return err
	}
	return nil
}

func UserDelete(user *Schema.User, userId string) (err error) {
	if err = Config.DB.Where("id = ?", userId).Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func UserFetchWithEmail(user *Schema.User, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}