package common

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// InitSession ...
func InitSession(r *http.Request, store *sessions.CookieStore) *sessions.Session {
	session, err := store.Get(r, "SessionId") // Don't ignore the error in real code
	LogFatal(err)

	if session.IsNew { //Set some cookie options
		session.Options.Domain = "localhost"
		session.Options.MaxAge = 60 * 60 * 60 * 24
		session.Options.HttpOnly = false
		session.Options.Secure = true
	}
	return session
}
