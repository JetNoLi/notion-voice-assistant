package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func CheckAuthorization(w *http.ResponseWriter, r *http.Request) {
	authCookie := ""

	for _, cookie := range r.Cookies() {
		if cookie.Name == "Authorization" {
			authCookie = cookie.Value
		}
	}

	if authCookie == "" {
		fmt.Println("No Cookie")
		*r = *r.Clone(context.WithValue(r.Context(), "Authorization", false))
		return
	}

	fmt.Printf("Cookie is %s", authCookie)
	*r = *r.Clone(context.WithValue(r.Context(), "Authorization", true))

}
