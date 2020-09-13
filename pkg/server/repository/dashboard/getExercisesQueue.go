package dashboardrepository

import (
	"context"
	"fmt"
	"mytrainingsserver/pkg/common"
	"mytrainingsserver/pkg/server/repository"
	"os"
)

// GetExercisesQueue ...
func GetExercisesQueue(trainingID int) []int {
	var exerciseArr []int
	var item int

	db := repository.Connect()

	rows, err := db.Query(context.Background(), fmt.Sprintf("Select id from exercises where training_id = %v ", trainingID))
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		err := rows.Scan(&item)
		common.LogFatal(err)

		exerciseArr = append(exerciseArr, item)
	}

	err = db.Close(context.Background())
	common.LogFatal(err)

	return exerciseArr
}
