package abstractfactory

//IUser interface
type IUser interface {
	GetUser() *User
}

//User struct
type User struct {
	Name               string
	Document           string
	Tipo               string
	ValidationNameDone bool
	ValidationTypeDone bool
}

//GetUser func
func (user *User) GetUser() *User {
	return user
}
