package tchat

// User represent system user
type User struct {
	ID       uint
	Nickname string `json:"nickname"`
	Login    string `json:"login"`
}

// UserService interface for user actions
type UserService interface {
	User(id uint) (*User, error)
	Users() ([]*User, error)
	Save(u *User) error
	Create(u *User) error
	Delete(id int) error
}
