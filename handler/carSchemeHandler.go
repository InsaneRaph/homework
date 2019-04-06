package handler

import (
	"encoding/json"
	"homeworkprojet/dto"
	"homeworkprojet/service"
	"homeworkprojet/utils"
	"net/http"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {

	cardSchemeDto := &dto.CardSchemeDto{}
	err := json.NewDecoder(r.Body).Decode(cardSchemeDto) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message("Invalid request"))
		return
	}

	schemeDto, err := service.GetCardSchemeService().Create(cardSchemeDto)
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message(err.Error()))
	} else {
		utils.Respond(w, http.StatusCreated, schemeDto)

	}
}

func GetUserCards(w http.ResponseWriter, r *http.Request) {

	cards := service.GetCardSchemeService().GetCardSchemesForUser(r.Context().Value("user").(uint))

	utils.Respond(w, http.StatusOK, cards)

}

func DeleteUserCards(w http.ResponseWriter, r *http.Request) {

	err := service.GetCardSchemeService().DeleteCardSchemeForUser(r.Context().Value("user").(uint),
		uint(utils.IntParam(r, "card")))
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
