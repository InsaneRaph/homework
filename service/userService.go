package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"homeworkprojet/config"
	"homeworkprojet/dto"
	"homeworkprojet/model"
	"strings"
)

type UserSvc interface {
	Create(userDto *dto.UserDto) (*dto.UserDto, error)
	Login(email, password string) (*dto.UserDto, error)
	GetUser(u uint) *dto.UserDto
}
type userSvc struct {
}

var uSvc *userSvc

func GetUserService() UserSvc {
	if uSvc == nil {
		uSvc = &userSvc{}
	}
	return uSvc
}

//Validate incoming user details...
func (svc *userSvc) validate(user *dto.UserDto) (bool, error) {

	if !strings.Contains(user.Email, "@") {
		return false, errors.New("email address is required")
	}

	if len(user.Password) < 6 {
		return false, errors.New("password is required")
	}

	//check for errors and duplicate emails
	temp, err := model.GetUserByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, errors.New("connection error. Please retry")

	}
	if temp.Email != "" {
		return false, errors.New("email address already in use by another user")
	}

	return true, nil
}

func (svc *userSvc) Create(userDto *dto.UserDto) (*dto.UserDto, error) {

	if ok, err := svc.validate(userDto); !ok {
		return nil, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	userDto.Password = string(hashedPassword)

	userModel := userDto.ToUser()
	err := userModel.Create()

	if err != nil {
		return nil, errors.New("failed to create userDto, connection error")
	}

	//Create new JWT token for the newly registered user

	return svc.createDtoWithToken(&userModel), nil
}

func (svc *userSvc) createDtoWithToken(user *model.User) *dto.UserDto {
	userDto := dto.FromUser(user)
	tk := &model.Token{UserId: userDto.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(config.AppConfig.JwtSecret))
	userDto.Token = tokenString
	userDto.Password = ""

	return userDto
}

func (svc *userSvc) Login(email, password string) (*dto.UserDto, error) {

	user, err := model.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return svc.createDtoWithToken(user), nil
}

func (svc *userSvc) GetUser(u uint) *dto.UserDto {

	user := dto.FromUser(model.GetUser(u))

	if user.Email == "" { //User not found!
		return nil
	}

	user.Password = ""
	return user
}
