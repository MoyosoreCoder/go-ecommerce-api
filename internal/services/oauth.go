package services
import (
	"golang.org/x/oauth2" 
	"golang.org/x/oauth2/google"
	"os"
)

func GetOauthConfig() *oauth2.Config {
	conf := &oauth2.Config{
		//client id is the application's ID
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_S ECRET"),
		RedirectURL: os.Getenv("CLIENT_REDIRECT"),
		Scopes: []string{"openid", "email", "profile"},
		Endpoint: google.Endpoint,
	}
	return conf
}