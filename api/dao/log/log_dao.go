package logdao

import (
	"github.com/fedehsq/api/db"
	"github.com/fedehsq/api/models/log"
)

func Insert(log *log.Log) error {
	_, err := db.DB.Exec("INSERT INTO logs (time, ip, caller_identity, method, route, body) VALUES ($1, $2, $3, $4, $5, $6)", log.Time, log.Ip, log.CallerIdentity, log.Method, log.Route, log.Body)
	if err != nil {
		return err
	}
	return nil
}

// func GetAll() ([]log.Log, error) {
// 	rows, err := db.DB.Query("SELECT * FROM logs")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 
// 	var logs []log.Log
// 	for rows.Next() {
// 		var log log.Log
// 		err := rows.Scan(&log.Id, &log.Time, &log.Ip, &log.CallerIdentity, &log.Method, &log.Route, &log.Body)
// 		if err != nil {
// 			return nil, err
// 		}
// 		logs = append(logs, log)
// 	}
// 	return logs, nil
// }
