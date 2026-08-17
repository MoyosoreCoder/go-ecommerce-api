package handlers

import (
	"net/http"

	"github.com/MoyosoreCoder/go-ecommerce-api/internal/services"
)

func  GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	conf := services.GetOauthConfig()
	
	url := conf.AuthCodeURL("state")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
