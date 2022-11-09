package user

type User struct {
	Id       int	`json:"id" example:"1"`
	Username string `json:"username" example:"user"`
	Password string `json:"password" example:"pwd"`
}
