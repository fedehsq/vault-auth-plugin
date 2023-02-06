package remotehostusersdao

import (
	"errors"
	"github.com/fedehsq/api/db"
	"github.com/fedehsq/api/models/remote-host-users"
)

func Get(remoteHostId int, userId int) (*remotehostusers.RemoteHostUsers, error) {
	var remoteHost remotehostusers.RemoteHostUsers
	err := db.DB.QueryRow("SELECT * FROM remote_host_users WHERE remote_host_id = $1 AND user_id = $2", remoteHostId, userId).Scan(&remoteHost.Id, &remoteHost.RemoteHostId, &remoteHost.UserId)
	if err != nil {
		return nil, err
	}
	return &remoteHost, nil
}

func GetAll(remoteHostId int) ([]*remotehostusers.RemoteHostUsers, error) {
	rows, err := db.DB.Query("SELECT * FROM remote_host_users where remote_host_id = $1", remoteHostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	remoteHostUsers := make([]*remotehostusers.RemoteHostUsers, 0)
	for rows.Next() {
		remoteHostUser := new(remotehostusers.RemoteHostUsers)
		err := rows.Scan(&remoteHostUser.Id, &remoteHostUser.RemoteHostId, &remoteHostUser.UserId)
		if err != nil {
			return nil, err
		}
		remoteHostUsers = append(remoteHostUsers, remoteHostUser)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return remoteHostUsers, nil
}

func Insert(remoteHostId int, userId int) error {
	_, err := db.DB.Exec("INSERT INTO remote_host_users (remote_host_id, user_id) VALUES ($1, $2)", remoteHostId, userId)
	if err != nil {
		return err
	}
	return nil
}

func Delete(remoteHostId int, userId int) error {
	res, err := db.DB.Exec("DELETE FROM remote_host_users WHERE remote_host_id = $1 AND user_id = $2", remoteHostId, userId)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("remoteHostUsers not found")
	}
	return nil
}
