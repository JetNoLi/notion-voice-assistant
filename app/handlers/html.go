package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/services"
	"github.com/jetnoli/notion-voice-assistant/utils"
	"github.com/jetnoli/notion-voice-assistant/view/pages/signup"
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

	cookie, err := utils.GenerateAuthCookie(user.Id)

	if err != nil {
		http.Error(w, "error creating cookie: "+err.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(w, cookie)

	err = signup.Success(user.Username).Render(r.Context(), w)

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

	cookie, err := utils.GenerateAuthCookie(user.Id)

	if err != nil {
		http.Error(w, "error creating cookie: "+err.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(w, cookie)
	w.Header().Set("HX-Redirect", "/")

	err = signup.Success(user.Username).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "error returning file: \n"+err.Error(), http.StatusInternalServerError)
	}

}

// func Takes Generic I & D
// Accepts :
//	- Service which Takes I as args
//  	- Returns D, error
//	- Function, which maps html data key -> html data, based on Generic D
