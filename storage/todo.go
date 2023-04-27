package storage

type Todo struct {
	ID   string
	Text string
	Done bool
	User *User
}

type User struct {
	ID   string
	Name string
}
