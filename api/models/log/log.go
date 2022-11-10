package log

import (
	"time"
)

type Log struct {
	Id      int `json:"id" example:"1"`
	Time    time.Time `json:"time" example:"2022-10-27 10:18:47.791249"`
	Ip      string `json:"ip" example:"127.0.0.1:50336"`
	Command string `json:"command" example:"Signin User"`
}

func (l *Log) String() string {
	return l.Time.Format(time.RFC3339) + " " + l.Ip + " " + l.Command
}
