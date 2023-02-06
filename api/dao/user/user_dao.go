package userdao

import (
	"errors"
	"github.com/fedehsq/api/db"
	"github.com/fedehsq/api/models/user"
)

func Get(id int) (*user.User, error) {
	// Query the database for the user
	var user user.User
	err := db.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetByUsername(username string) (*user.User, error) {
	// Query the database for the user
	var user user.User
	err := db.DB.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Insert(user *user.User) error {
	_, err := db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func Delete(username string) error {
	res, err := db.DB.Exec("DELETE FROM users WHERE username = $1", username)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func Update(username string, password string) (*user.User, error) {
	_, err := db.DB.Exec("UPDATE users SET password = $1 WHERE username = $2", password, username)
	if err != nil {
		return nil, err
	}
	user, err := GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAll() ([]user.User, error) {
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var user user.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
