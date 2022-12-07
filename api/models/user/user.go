package user

type User struct {
	Id       int
	Username string `json:"username" example:"user"`
	Password string `json:"password" example:"pwd"`
}

func (u *User) Json() string {
	return "{\"username\":\"" + u.Username + "\"}"
}
