package routes

import (
	"encoding/json"
	"mytrainingsserver/pkg/common"
	dashboardrepository "mytrainingsserver/pkg/server/repository/dashboard"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// GetExercises ....
func (d *Dashboard) GetExercises() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		intParam, err := strconv.Atoi(param["id"])
		result := "Неверный параметр"
		if err != nil {
			common.LogFatal(err)
			json.NewEncoder(w).Encode(result)
		} else {
			result := dashboardrepository.GetExercises(intParam)

			json.NewEncoder(w).Encode(result)
		}

	}
}

// GetExercisesQueue ....
func (d *Dashboard) GetExercisesQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		intParam, err := strconv.Atoi(param["id"])
		result := "Неверный параметр"
		if err != nil {
			common.LogFatal(err)
			json.NewEncoder(w).Encode(result)
		} else {
			result := dashboardrepository.GetExercisesQueue(intParam)

			json.NewEncoder(w).Encode(result)
		}

	}
}

// GetExercise ....
func (d *Dashboard) GetExercise() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		var exerciseID map[string]int

		json.NewDecoder(r.Body).Decode(&exerciseID)

		println(exerciseID)

		intParam, err := strconv.Atoi(param["id"])
		result := "Неверный параметр"
		if err != nil {
			common.LogFatal(err)

			json.NewEncoder(w).Encode(result)
		} else {
			result := dashboardrepository.GetExercise(intParam, exerciseID["id"])

			json.NewEncoder(w).Encode(result)
		}

	}
}
