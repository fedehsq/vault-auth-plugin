package admindao

import (
	"github.com/fedehsq/api/db"
	"github.com/fedehsq/api/models/admin"
)

func GetByUsername(username string) (*admin.Admin, error) {
	// Query the database for the admin
	var admin admin.Admin
	err := db.DB.QueryRow("SELECT * FROM admins WHERE username = $1", username).Scan(&admin.Id, &admin.Username, &admin.Password)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
