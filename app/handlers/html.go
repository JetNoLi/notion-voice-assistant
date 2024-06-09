package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/services"
	"github.com/jetnoli/notion-voice-assistant/utils"
	"github.com/jetnoli/notion-voice-assistant/wrappers/serve"
)

// TODO: Allow a way to specify which routes are restricted when serving static html

func SignUpHtmx(w http.ResponseWriter, r *http.Request) {

	userDetails := &services.SignUpRequestBody{}

	err := json.NewDecoder(r.Body).Decode(&userDetails)

	if err != nil {
		http.Error(w, "cannot read json: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.SignUp(userDetails)

	if err != nil {
		http.Error(w, "error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	htmlData := make(map[string]string)

	htmlData["username"] = user.Username

	html, err := serve.AndInjectHtml(("static/html/responses/signup-success.html"), htmlData)

	if err != nil {
		http.Error(w, "error Reading file: \n"+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(html)

	if err != nil {
		http.Error(w, "error returning file: \n"+err.Error(), http.StatusInternalServerError)
	}
}

func SignInHtmx(w http.ResponseWriter, r *http.Request) {

	userDetails := &services.SignInRequestBody{}

	err := json.NewDecoder(r.Body).Decode(&userDetails)

	if err != nil {
		http.Error(w, "cannot read json: "+err.Error(), http.StatusBadRequest)
	}

	user, err := services.SignIn(userDetails)

	if err != nil {
		http.Error(w, "error authenticating in user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	htmlData := make(map[string]string)

	htmlData["username"] = user.Username

	html, err := serve.AndInjectHtml(("static/html/responses/signup-success.html"), htmlData)

	if err != nil {
		http.Error(w, "error reading file:\n"+err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO: Implement Based on JWT
	http.SetCookie(w, utils.GenerateAuthCookie("true"))

	_, err = w.Write(html)

	if err != nil {
		http.Error(w, "error returning file: \n"+err.Error(), http.StatusInternalServerError)
	}
}

// func Takes Generic I & D
// Accepts :
//	- Service which Takes I as args
//  	- Returns D, error
//	- Function, which maps html data key -> html data, based on Generic D
