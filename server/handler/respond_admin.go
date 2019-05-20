package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/sentriz/gonic/model"
)

type templateData struct {
	Flashes                []interface{}
	User                   *model.User
	SelectedUser           *model.User
	AllUsers               []*model.User
	ArtistCount            int
	AlbumCount             int
	TrackCount             int
	CurrentLastFMAPIKey    string
	CurrentLastFMAPISecret string
	RequestRoot            string
}

func renderTemplate(w http.ResponseWriter, r *http.Request,
	tmpl *template.Template, data *templateData) {
	session := r.Context().Value(contextSessionKey).(*sessions.Session)
	if data == nil {
		data = &templateData{}
	}
	data.Flashes = session.Flashes()
	session.Save(r, w)
	user, ok := r.Context().Value(contextUserKey).(*model.User)
	if ok {
		data.User = user
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("500 when executing: %v", err), 500)
		return
	}
}