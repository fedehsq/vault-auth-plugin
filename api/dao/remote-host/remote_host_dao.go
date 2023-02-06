package remotehostdao

import (
	"errors"

	"github.com/fedehsq/api/db"
	remotehost "github.com/fedehsq/api/models/remote-host"
)

func GetByIp(ip string) (*remotehost.RemoteHost, error) {
	// Query the database for the remoteHost
	var remoteHost remotehost.RemoteHost
	err := db.DB.QueryRow("SELECT * FROM remote_hosts WHERE ip = $1", ip).Scan(&remoteHost.Id, &remoteHost.Ip)
	if err != nil {
		return nil, err
	}
	return &remoteHost, nil
}

func Insert(remoteHost *remotehost.RemoteHost) error {
	_, err := db.DB.Exec("INSERT INTO remote_hosts (ip) VALUES ($1)", remoteHost.Ip)
	if err != nil {
		return err
	}
	return nil
}

func Delete(ip string) error {
	res, err := db.DB.Exec("DELETE FROM remote_hosts WHERE ip = $1", ip)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("remoteHost not found")
	}
	return nil
}
