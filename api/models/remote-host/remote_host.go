package remotehost

type RemoteHost struct {
	Id int `json:"id"`
	Ip string `json:"ip"`
}

func (r *RemoteHost) Json() string {
	return "{\"ip\":\"" + r.Ip + "\"}"
}