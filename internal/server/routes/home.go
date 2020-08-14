package routes

import (
	"encoding/json"
	"log"
	"mytrainingsserver/pkg/common"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
)

// Home ...
type Home struct{}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_AUTH_KEY")))

// var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))

// GetHome ...
func (h Home) GetHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)

	}
}

// GetLoginHome ...
func (h Home) GetLoginHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := "Login"
		println("Login")
		json.NewEncoder(w).Encode(result)
	}
}

// GetRegistrationHome ...
func (h Home) GetRegistrationHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var user models.User

		// result := "Registration"
		// json.NewEncoder(w).Encode(result)
	}
}

// GetRestorePassHome ...
func (h Home) GetRestorePassHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := "RestorePass"
		json.NewEncoder(w).Encode(result)
	}
}

// PostLoginHome ...
func (h Home) PostLoginHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Login struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}

		var user Login

		json.NewDecoder(r.Body).Decode(&user)

		session, err := store.Get(r, "sessionId")
		common.LogFatal(err)

		session.Values["userId"] = "15"
		session.Values["userName"] = "John"

		err = session.Save(r, w)
		common.LogFatal(err)

		// store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte(os.Getenv("SESSION_KEY")))

		log.Println(user)

		json.NewEncoder(w).Encode(user)

		// session := common.InitSession(r, store)
		// session.Values["userId"] = "15"
		// err := sessions.Save(r, w)
		// common.LogFatal(err)
	}
}

// PostRegistrationHome ...
func (h Home) PostRegistrationHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Register struct {
			Login       string `json:"reglogin"`
			Password    string `json:"regpassword"`
			PassConfirm string `json:"regPassConfirm"`
		}

		var user Register

		json.NewDecoder(r.Body).Decode(&user)

		cookie := http.Cookie{
			Name:    "cookieID",
			Value:   "myCookieID",
			Expires: time.Now().Add(20 * time.Hour),
		}

		log.Println(user)

		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)

		// json.NewEncoder(w).Encode(&User)
		// res.Body = ioutil.NopCloser()
	}
}

// PostRestorePassHome ...
func (h Home) PostRestorePassHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := "Post RestorePass"
		json.NewEncoder(w).Encode(result)
	}
}
