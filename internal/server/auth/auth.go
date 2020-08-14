package auth

import (
	"log"
	"time"

	"github.com/adam-hanna/jwt-auth/jwt"
)

// Auth ...
type Auth struct {
	SecuredRoute jwt.Auth
}

// New ...
func (a Auth) New(auth jwt.Auth) jwt.Auth {

	authErr := jwt.New(&auth, jwt.Options{
		SigningMethodString:   "RS256",
		PrivateKeyLocation:    "keys/app.rsa",
		PublicKeyLocation:     "keys/app.rsa.pub",
		RefreshTokenValidTime: 72 * time.Hour,
		AuthTokenValidTime:    15 * time.Minute,
		Debug:                 false,
		IsDevEnv:              true,
	})
	if authErr != nil {
		log.Println("Error initializing the JWT's!")
		log.Fatal(authErr)
	}

	return auth
}
