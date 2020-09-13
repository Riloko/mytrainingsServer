package dashboardrepository

import (
	"context"
	"fmt"
	"mytrainingsserver/internal/server/models"
	"mytrainingsserver/pkg/common"
	"mytrainingsserver/pkg/server/repository"
	"time"
)

// GetExercise ...
func GetExercise(trainingID int, exerciseID int) map[string]interface{} {
	var item map[string]interface{}
	var exercise models.Exercise

	println(trainingID)
	println(exerciseID)

	db := repository.Connect()

	db.QueryRow(context.Background(), fmt.Sprintf("Select * from exercises where training_id = %v and id = %v ", trainingID, exerciseID)).Scan(&exercise.ID, &exercise.TrainingID, &exercise.Image, &exercise.Name, &exercise.Description, &exercise.Type, &exercise.IsGeo, &exercise.Difficulty, &exercise.Time)

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

	err := db.Close(context.Background())
	common.LogFatal(err)

	return item
}
