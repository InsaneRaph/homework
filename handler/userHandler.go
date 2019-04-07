package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"homeworkprojet/dto"
	"homeworkprojet/service"
	"homeworkprojet/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &dto.UserDto{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message("Invalid request"))
		return
	}

	usr, err := service.GetUserService().Create(user) //Create user
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message(err.Error()))
	} else {
		utils.Respond(w, http.StatusCreated, usr)

	}
}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	user := &dto.UserDto{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message("Invalid request"))
		return
	}

	usr, err := service.GetUserService().Login(user.Email, user.Password)
	status := http.StatusOK

	if err != nil {
		msg := err.Error()
		status = http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			msg = "Email address not found"
			status = http.StatusNotFound
		} else if err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			status = http.StatusForbidden
			msg = "Invalid login credentials. Please try again"
		}
		utils.Respond(w, status, utils.Message(msg))
		return
	}
	utils.Respond(w, status, usr)
}

func UserInfo(w http.ResponseWriter, r *http.Request) {

	usr := service.GetUserService().GetUser(getCurrentUser(r)) //Create user
	if usr != nil {
		utils.Respond(w, http.StatusOK, usr)
	} else {
		utils.Respond(w, http.StatusNotFound, utils.Message("user not found"))
	}
}

func getCurrentUser(r *http.Request) uint {
	return r.Context().Value("user").(uint)
}
