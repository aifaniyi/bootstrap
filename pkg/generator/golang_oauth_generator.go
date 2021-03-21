package generator

import (
	"github.com/aifaniyi/bootstrap/templates/golang/web/auth/oauth"
)

func generateRandomStateString() string {
	return oauth.RandomStateString
}

func generateOauthConfig(auth string) string {
	return oauth.Configs[auth]
}

func generateOauthHandler(auth string) string {
	return oauth.Handlers[auth]
}
