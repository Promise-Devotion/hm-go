package models

type User struct {
	UserId    int    `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	NickName  string `json:"nickname"`
	Email     string `json:"email"`
	Sex       int    `json:"sex"`
}

func (User) TableName() string {
	return "t_user"
}
