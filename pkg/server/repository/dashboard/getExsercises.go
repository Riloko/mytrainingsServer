package dashboardrepository

import (
	"context"
	"fmt"
	"mytrainingsserver/pkg/common"
	"mytrainingsserver/pkg/server/repository"
	"os"
	"time"

	"mytrainingsserver/internal/server/models"
)

// GetExercises ...
func GetExercises(trainingID int) []map[string]interface{} {
	var exerciseArr []map[string]interface{}
	var item map[string]interface{}
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
		t := time.Time(exercise.Time)

		item = map[string]interface{}{
			"id":          exercise.ID,
			"training_id": exercise.TrainingID,
			"image_url":   exercise.Image,
			"name":        exercise.Name,
			"description": exercise.Description,
			"type":        exercise.Type,
			"geo":         exercise.IsGeo,
			"difficulty":  exercise.Difficulty,
			"time":        t.Minute(),
		}

		exerciseArr = append(exerciseArr, item)
	}

	err = db.Close(context.Background())
	common.LogFatal(err)

	return exerciseArr
}
