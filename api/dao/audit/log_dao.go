package auditdao

import (
	"fmt"
	sqldb "vault-auth-plugin/api/db"
	"vault-auth-plugin/api/models/audit"
)

func Insert(log *audit.Log) error {
	_, err := sqldb.DB.Exec("INSERT INTO logs (time, ip, command) VALUES ($1, $2, $3)", log.Time, log.Ip, log.Command)
	if err != nil {
		return err
	}
	return nil
}

func GetAll() ([]audit.Log, error) {
	rows, err := sqldb.DB.Query("SELECT * FROM logs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []audit.Log
	for rows.Next() {
		var log audit.Log
		err := rows.Scan(&log.Id, &log.Time, &log.Ip, &log.Command)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	fmt.Println("Inserting log: ", logs)

	return logs, nil
}
