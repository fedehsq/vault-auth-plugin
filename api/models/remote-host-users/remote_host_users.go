package remotehostusers

type RemoteHostUsers struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	RemoteHostId string `json:"remote_host_id"`
}