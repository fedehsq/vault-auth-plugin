package dao

import (
	sqldb "vault-auth-plugin/server/db"
	"vault-auth-plugin/server/models"
)

func GetUserByUsername(username string) (*models.User, error) {
	// Query the database for the user
	var user models.User
	err := sqldb.DB.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func InsertUser(user *models.User) error {
	_, err := sqldb.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]models.User, error) {
	rows, err := sqldb.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
