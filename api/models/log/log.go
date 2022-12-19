package log

type Log struct {
	Id             int    `json:"id" example:"1"`
	Time           string `json:"time" example:"Tue Nov 10 23:00:00 UTC 2009"`
	Ip             string `json:"ip" example:"127.0.0.1:50336"`
	CallerIdentity string `json:"callerIdentity" example:"admin"`
	Method         string `json:"method" example:"POST"`
	Route          string `json:"route" example:"/api/v1/admin/signin"`
	Body           string `json:"body" example:"{username: admin, password: ********}"`
}

func (l *Log) String() string {
	return l.Time + " " + l.Ip + " " + l.CallerIdentity + " " + l.Method + " " + l.Route + " " + l.Body
}
