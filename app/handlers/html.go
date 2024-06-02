package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config/client"
	"github.com/jetnoli/notion-voice-assistant/models/user"
	"github.com/jetnoli/notion-voice-assistant/services"
	"github.com/jetnoli/notion-voice-assistant/wrappers/fetch"
	"github.com/jetnoli/notion-voice-assistant/wrappers/serve"
)

//TODO: Automate all folders in html folder get served

func ServeRoot(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	res, err := client.WhisperApi.Get("/", fetch.ApiGetRequestOptions{})

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(502)
		return
	}

	defer res.Body.Close()

	html, err := serve.Html("static/html/index.html")

	if err != nil {
		http.Error(w, "Error Reading file:\n"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(html)
}

func ServeSignUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	html, err := serve.Html("static/html/signup.html")

	if err != nil {
		http.Error(w, "Error Reading file:\n"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(html)
}

func SignupHtmx(w http.ResponseWriter, r *http.Request) {

	userDetails := &user.User{}

	err := json.NewDecoder(r.Body).Decode(&userDetails)

	if err != nil {
		http.Error(w, "cannot read json: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.SignUpUser(userDetails)

	if err != nil {
		http.Error(w, "error creating user: "+err.Error(), http.StatusInternalServerError)
	}

	htmlData := make(map[string]string)

	htmlData["username"] = user.Username

	html, err := serve.AndInjectHtml(("static/html/responses/signup-success.html"), htmlData)

	if err != nil {
		http.Error(w, "Error Reading file:\n"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(html)
}

// func Takes Generic I & D
// Accepts :
//	- Service which Takes I as args
//  	- Returns D, error
//	- Function, which maps html data key -> html data, based on Generic D
