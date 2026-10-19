// this codes handles google call back
package handlers

import (
	"fmt"
	"net/http"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/services"
)
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	conf := services.GetOauthConfig()
	code := r.URL.Query().Get(("code"))


	//tok, err := conf.Exchange(r.Context() , code)
	tok, err := conf.Exchange(r.Context(), code)
	if err != nil {
 		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = tok

	fmt.Fprintln(w, "Google login successful!")

	client := conf.Client(r.Context(), tok)
 
	
	_ = client


	//user's profile
	resp, err := client.Get("https://openidconnect.googleapis.com/v1/userinfo")
	if err != nil {
		http.Error(w, "Failed to retrieve Google profile", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
}