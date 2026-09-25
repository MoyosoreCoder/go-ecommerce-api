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


	//tok, err := conf.Exchange(r.Context(), code)
	tok, err := conf.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = tok

	fmt.Fprintln(w, "Google login successful!")

}