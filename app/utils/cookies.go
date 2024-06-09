package utils

import "net/http"

func GenerateAuthCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		MaxAge:   60 * 60 * 24 * 7, // 1 Week
		Path:     "/",
	}
}
