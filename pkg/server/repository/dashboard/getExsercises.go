package dashboardrepository

import (
	"context"
	"fmt"
	"mytrainingsserver/pkg/common"
	"mytrainingsserver/pkg/server/repository"
	"os"

	"mytrainingsserver/internal/server/models"
)

// GetExercises ...
func GetExercises(trainingID int) []models.Exercise {
	var exercises []models.Exercise
	var exercise models.Exercise

	db := repository.Connect()

	rows, err := db.Query(context.Background(), fmt.Sprintf("Select * from exercises where training_id = %v ", trainingID))
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		err := rows.Scan(&exercise.ID, &exercise.TrainingID, &exercise.Image, &exercise.Name, &exercise.Description, &exercise.Type, &exercise.IsGeo, &exercise.Difficulty, &exercise.Time)
		common.LogFatal(err)

		exercises = append(exercises, exercise)
	}

	err = db.Close(context.Background())
	common.LogFatal(err)

	return exercises
}
