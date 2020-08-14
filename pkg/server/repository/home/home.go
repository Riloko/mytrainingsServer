package homerepository

import (
	"mytrainingsserver/pkg/common"
	"mytrainingsserver/pkg/server/repository"

	"mytrainingsserver/internal/server/models"
)

// GetUsers ...
func GetUsers() []models.User {
	var users []models.User
	var user models.User

	db := repository.Connect()

	rows, err := db.Query("Select * from users")
	common.LogFatal(err)

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Login, &user.Pass, &user.Rank, &user.Score)
		common.LogFatal(err)

		users = append(users, user)
	}

	err = db.Close()
	common.LogFatal(err)

	return users
}
