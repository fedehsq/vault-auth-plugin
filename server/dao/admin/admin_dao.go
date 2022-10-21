package admindao

import (
	sqldb "vault-auth-plugin/server/db"
	"vault-auth-plugin/server/models/admin"
)

func GetByUsername(username string) (*admin.Admin, error) {
	// Query the database for the admin
	var admin admin.Admin
	err := sqldb.DB.QueryRow("SELECT * FROM admins WHERE username = $1", username).Scan(&admin.Id, &admin.Username, &admin.Password)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

