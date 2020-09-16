package abstractfactory

//IUser interface
type IUser interface {
	SetName(name string)
	SetDocument(document string)
	SetTipo(tipo string)
	GetName() string
	GetDocument() string
	GetTipo() string
}

//User struct
type User struct {
	name     string
	document string
	tipo     string
}

//SetName set
func (user *User) SetName(name string) {
	user.name = name
}

//GetName get
func (user *User) GetName() string {
	return user.name
}

//SetDocument set
func (user *User) SetDocument(document string) {
	user.document = document
}

//GetDocument get
func (user *User) GetDocument() string {
	return user.document
}

//SetTipo set
func (user *User) SetTipo(tipo string) {
	user.tipo = tipo
}

//GetTipo get
func (user *User) GetTipo() string {
	return user.tipo
}
