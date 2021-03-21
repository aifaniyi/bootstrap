package oauth

// Configs :
var Configs map[string]string

// vars
const (
	RandomStateString = `
	package oauth

// getRandomStateString :
func getRandomStateString() string {
	return "pseudo-random"
}
`
)

func init() {
	Configs = map[string]string{
		"google": `
		package oauth
	
	import (
		"gitlab.com/aifaniyi/go-libs/conf"
		"golang.org/x/oauth2"
		"golang.org/x/oauth2/google"
	)
	
	var googleOauthConfig *oauth2.Config
	
	func init() {
		googleOauthConfig = &oauth2.Config{
			RedirectURL:  conf.GetStringEnv("CALLBACK_URL", "http://localhost:8080/google/callback"),
			ClientID:     conf.GetStringEnv("GOOGLE_CLIENT_ID", "google_client_id"),
			ClientSecret: conf.GetStringEnv("GOOGLE_CLIENT_SECRET", "google_client_secret"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		}
	}
	
	// getGoogleOauthConfig :
	func getGoogleOauthConfig() *oauth2.Config {
		return googleOauthConfig
	}
	`,
		"facebook": `
	package oauth

import (
	"gitlab.com/aifaniyi/go-libs/conf"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var facebookOauthConfig *oauth2.Config

func init() {
	facebookOauthConfig = &oauth2.Config{
		RedirectURL:  conf.GetStringEnv("CALLBACK_URL", "http://localhost:8080/facebook/callback"),
		ClientID:     conf.GetStringEnv("FACEBOOK_CLIENT_ID", "facebook_client_id"),
		ClientSecret: conf.GetStringEnv("FACEBOOK_CLIENT_SECRET", "facebook_client_secret"),
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
}

// getFacebookOauthConfig :
func getFacebookOauthConfig() *oauth2.Config {
	return facebookOauthConfig
}
`,
	}
}
