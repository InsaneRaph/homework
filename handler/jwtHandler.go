package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"homeworkprojet/config"
	models "homeworkprojet/model"
	"homeworkprojet/utils"
	"net/http"
	"strings"
)

func JwtAuthentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = utils.Message("Missing auth token")
			utils.Respond(w, http.StatusForbidden, response)
			return
		}

		splited := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splited) != 2 {
			response = utils.Message("Invalid/Malformed auth token")

			utils.Respond(w,http.StatusForbidden, response)
			return
		}

		tokenPart := splited[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JwtSecret), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = utils.Message("Malformed authentication token")
			utils.Respond(w,http.StatusForbidden, response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = utils.Message("Token is not valid.")
			utils.Respond(w,http.StatusForbidden, response)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
