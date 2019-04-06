package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

var userTable = "users"

//a struct to rep user account
type User struct {
	gorm.Model
	Email    string `gorm:"unique_index"`
	Password string
}

func GetUserByEmail(email string) (*User, error) {

	temp := &User{}
	err := DB.Table(userTable).Where("email = ?", email).First(temp).Error
	return  temp, err

}

func (user *User) Create() error {
	return DB.Create(user).Error
}

func GetUser(u uint) *User {

	user := &User{}
	DB.Table(userTable).Where("id = ?", u).First(user)
	return user
}
