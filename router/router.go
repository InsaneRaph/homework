package router

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"homeworkprojet/handler"
)

var router *mux.Router

func RegisterRoutes() *mux.Router {

	router = mux.NewRouter()
	httpHandlers := alice.New(handler.JwtAuthentication)

	router.HandleFunc("/api/user/login", handler.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/new", handler.CreateUser).Methods("POST")

	router.Handle("/api/user", httpHandlers.ThenFunc(handler.UserInfo)).Methods("GET")
	router.Handle("/api/cards", httpHandlers.ThenFunc(handler.GetUserCards)).Methods("GET")
	router.Handle("/api/cards", httpHandlers.ThenFunc(handler.CreateCard)).Methods("POST")
	router.Handle("/api/cards/{card}", httpHandlers.ThenFunc(handler.DeleteUserCards)).Methods("DELETE")
	router.Handle("/api/cards/{card}/payments", httpHandlers.ThenFunc(handler.GetPayments)).Methods("GET")

	return router
}
