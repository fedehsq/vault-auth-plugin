package logdao

import (
	"fmt"
	"github.com/fedehsq/api/db"
	"github.com/fedehsq/api/models/log"
)

func Insert(log *log.Log) error {
	_, err := db.DB.Exec("INSERT INTO logs (time, ip, command) VALUES ($1, $2, $3)", log.Time, log.Ip, log.Command)
	if err != nil {
		return err
	}
	return nil
}

func GetAll() ([]log.Log, error) {
	rows, err := db.DB.Query("SELECT * FROM logs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []log.Log
	for rows.Next() {
		var log log.Log
		err := rows.Scan(&log.Id, &log.Time, &log.Ip, &log.Command)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	fmt.Println("Inserting log: ", logs)

	return logs, nil
}
