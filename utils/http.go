package utils

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

func Respond(w http.ResponseWriter, httpStatus int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	err := json.NewEncoder(w).Encode(data)
	PanicOnError(err)
}

func IntParam(r *http.Request, param string) int {
	params := context.Get(r, "params").(httprouter.Params)
	id := params.ByName(param)
	if id != "" {
		p, _ := strconv.Atoi(id)
		return p
	}
	return -1
}

func StrParam(r *http.Request, param string) string {
	params := context.Get(r, "params").(httprouter.Params)
	return params.ByName(param)
}
