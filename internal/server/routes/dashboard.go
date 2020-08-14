package routes

import (
	"encoding/json"
	homerepository "mytrainingsserver/pkg/server/repository/home"
	"net/http"
)

// Dashboard ...
type Dashboard struct{}

// GetDashboard ....
func (d Dashboard) GetDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := homerepository.GetUsers()

		json.NewEncoder(w).Encode(result)
	}
}
