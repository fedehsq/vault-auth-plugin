package log

import (
	"time"
)

type Log struct {
	Id             int       `json:"id" example:"1"`
	Time           time.Time `json:"time" example:"2022-10-27 10:18:47.791249"`
	Ip             string    `json:"ip" example:"127.0.0.1:50336"`
	CallerIdentity string    `json:"callerIdentity" example:"admin"`
	Method         string    `json:"method" example:"POST"`
	Route          string    `json:"route" example:"/api/v1/admin/signin"`
	Body           string    `json:"body" example:"{username: admin, password: ********}"`
}

func (l *Log) String() string {
	return l.Time.String() + " " + l.Ip + " " + l.CallerIdentity + " " + l.Method + " " + l.Route + " " + l.Body
}
