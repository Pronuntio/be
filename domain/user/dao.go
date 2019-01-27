package user

type Dao interface {
	GetUser(ID uint64) (*User, error)
	DeleteUser(ID uint64) error
	CreateUser(user *User) error
	UpdateUser(user *User) error
}
