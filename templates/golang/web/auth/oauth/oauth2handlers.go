package oauth

// Handlers :
var Handlers map[string]string

func init() {
	Handlers = map[string]string{
		"google": `
		package oauth

		func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
			url := getGoogleOauthConfig().AuthCodeURL(getRandomStateString())
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		}
		
		func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
			content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
			if err != nil {
				fmt.Println(err.Error())
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			fmt.Fprintf(w, "Content: %s\n", content)
		}
		
		func getUserInfo(state string, code string) ([]byte, error) {
			config := getGoogleOauthConfig()
			stateString := getRandomStateString()
		
			if state != stateString {
				return nil, fmt.Errorf("invalid oauth state")
			}
			token, err := config.Exchange(oauth2.NoContext, code)
			if err != nil {
				return nil, fmt.Errorf("code exchange failed: %s", err.Error())
			}
		
			response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
			if err != nil {
				return nil, fmt.Errorf("failed getting user info: %s", err.Error())
			}
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return nil, fmt.Errorf("failed reading response body: %s", err.Error())
			}
			return contents, nil
		}
			`,
		"facebook": `
		package oauth
		
		func HandleFacebookLogin(w http.ResponseWriter, r *http.Request) {
			url := getFacebookOauthConfig().AuthCodeURL(getRandomStateString())
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		}
		
		func HandleFacebookCallback(w http.ResponseWriter, r *http.Request) {
			content, err := getFacebookInfo(r.FormValue("state"), r.FormValue("code"))
			if err != nil {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			fmt.Fprintf(w, "Content: %s\n", content)
		}
		
		func getFacebookInfo(state string, code string) ([]byte, error) {
			config := getFacebookOauthConfig()
			stateString := getRandomStateString()
		
			if state != stateString {
				return nil, fmt.Errorf("invalid oauth state")
			}
			token, err := config.Exchange(oauth2.NoContext, code)
			if err != nil {
				return nil, fmt.Errorf("code exchange failed: %s", err.Error())
			}
		
			response, err := http.Get("https://graph.facebook.com/me?fields=id,name,email&access_token=" + url.QueryEscape(token.AccessToken))
			if err != nil {
				return nil, fmt.Errorf("failed getting user info: %s", err.Error())
			}
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return nil, fmt.Errorf("failed reading response body: %s", err.Error())
			}
			return contents, nil
		}
		
		`,
	}
}
