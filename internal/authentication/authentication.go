package authentication

import (
	"net/http"

	"catalyst.api/config"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

const (
	// TODO change this key
	key          = "randomString"
	MaxAge       = 86400 * 30
	IsProduction = false
)

func NewAuthentication(cfg *config.Config) {
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProduction
	store.Options.SameSite = http.SameSiteLaxMode

	gothic.Store = store

	goth.UseProviders(
		github.New(cfg.AuthenticationConfig.GithubClientID, cfg.AuthenticationConfig.GithubClientSecret, "http://localhost:42069/auth/github/callback", "user", "repo"),
	)
}
