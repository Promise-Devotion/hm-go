package models

type User struct {
	Userid    int
	Firstname string
	Lastname  string
	Nickname  string
	Email     string
	sex       int
}

func (User) Tablename() string {
	return "t_user"
}
