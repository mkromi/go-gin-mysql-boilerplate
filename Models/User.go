package Models
import (
	_ "github.com/go-sql-driver/mysql"
	"go-gin-mysql-boilerplate/Config"
	"go-gin-mysql-boilerplate/Models/Schema"
)

func UserCreate(user *Schema.User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}