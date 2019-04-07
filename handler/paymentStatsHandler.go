package handler

import (
	"homeworkprojet/service"
	"homeworkprojet/utils"
	"net/http"
)

func GetPayments(w http.ResponseWriter, r *http.Request) {
	
	payments := service.GetPaymentStatsService().GetPayments(getCurrentUser(r),
		uint(utils.IntParam(r, "card")))
	utils.Respond(w, http.StatusOK, payments)
}
