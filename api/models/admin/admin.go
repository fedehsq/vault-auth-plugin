package admin

type Admin struct {
	Id       int    `json:"id"`
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password"`
}
