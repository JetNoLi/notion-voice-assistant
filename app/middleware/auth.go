package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/utils"
)

func CheckAuthorization(w *http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId")

	if userId == "" || userId == nil {
		http.Error(*w, "no access token provided", http.StatusUnauthorized)
		_ = r.Context().Done() //TODO: Double Check Works
	}
}

func DecodeToken(w *http.ResponseWriter, r *http.Request) {
	token := ""

	for _, cookie := range r.Cookies() {
		if cookie.Name == "Authorization" {
			token = cookie.Value
		}
	}

	if token == "" {
		fmt.Println("No Cookie")
		return
	}

	userId, err := utils.DecodeJwt(token)

	if err != nil {
		fmt.Println("error decoding token: " + err.Error())
		return
	}

	fmt.Println("userId ", userId)

	*r = *r.Clone(context.WithValue(r.Context(), "userId", userId))
}
