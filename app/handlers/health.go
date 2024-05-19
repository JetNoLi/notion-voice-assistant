package handlers

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("It's Alive"))
	defer r.Body.Close()
}
