package dashboardrepository

import (
	"context"
	"fmt"
	"mytrainingsserver/pkg/common"
	"mytrainingsserver/pkg/server/repository"
	"os"

	"mytrainingsserver/internal/server/models"
)

// GetTrainings ...
func GetTrainings(lim int, off int) []models.Training {
	var trainings []models.Training
	var training models.Training

	db := repository.Connect()

	rows, err := db.Query(context.Background(), fmt.Sprintf("Select * from Trainings limit %v offset %v", lim, off))
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		err := rows.Scan(&training.ID, &training.Name, &training.Description, &training.Difficulty, &training.Time, &training.Exp)
		common.LogFatal(err)

		trainings = append(trainings, training)
	}

	err = db.Close(context.Background())
	common.LogFatal(err)

	return trainings
}

// GetLimit ...
func GetLimit() int {
	var limit int

	db := repository.Connect()

	row := db.QueryRow(context.Background(), "Select count(*) from Trainings")
	row.Scan(&limit)

	err := db.Close(context.Background())
	common.LogFatal(err)

	return limit
}
