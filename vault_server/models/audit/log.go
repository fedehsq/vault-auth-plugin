package audit

import (
	"time"
)

type Log struct {
	Id      int
	Time    time.Time
	Ip      string
	Command string
}

func (l *Log) String() string {
	return l.Time.Format(time.RFC3339) + " " + l.Ip + " " + l.Command
}
