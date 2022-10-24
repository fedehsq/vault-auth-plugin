package userdao

import (
	"errors"
	sqldb "vault-auth-plugin/vault_server/db"
	"vault-auth-plugin/vault_server/models/user"
)

func GetByUsername(username string) (*user.User, error) {
	// Query the database for the user
	var user user.User
	err := sqldb.DB.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Insert(user *user.User) error {
	_, err := sqldb.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func Delete(username string) error {
	res, err := sqldb.DB.Exec("DELETE FROM users WHERE username = $1", username)
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
	_, err := sqldb.DB.Exec("UPDATE users SET password = $1 WHERE username = $2", password, username)
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
	rows, err := sqldb.DB.Query("SELECT * FROM users")
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
