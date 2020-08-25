package routes

import (
	"encoding/json"
	dashboardrepository "mytrainingsserver/pkg/server/repository/dashboard"
	"net/http"
)

// Dashboard ...
type Dashboard struct{}

// GetDashboard ....
func (d *Dashboard) GetDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// result := homerepository.GetUsers()

		// json.NewEncoder(w).Encode(result)
	}
}

// GetDashboardTrainings ....
func (d *Dashboard) GetDashboardTrainings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Params struct {
			Lim int `json:"limit"`
			Off int `json:"offset"`
		}
		var params Params

		json.NewDecoder(r.Body).Decode(&params)

		result := dashboardrepository.GetTrainings(params.Lim, params.Off)

		json.NewEncoder(w).Encode(result)
	}
}

// GetTrainingsLimit ....
func (d *Dashboard) GetTrainingsLimit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := dashboardrepository.GetLimit()

		json.NewEncoder(w).Encode(result)
	}
}
