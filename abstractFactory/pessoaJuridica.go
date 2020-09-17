package abstractfactory

//PessoaJuridica struct
type PessoaJuridica struct {
}

//MakeUser implementation of AbstractFactory
func (pessoa *PessoaJuridica) MakeUser(name string, document string, tipo string) IUser {
	return &User{
		Name:               name,
		Document:           document,
		Tipo:               tipo,
		ValidationNameDone: false,
		ValidationTypeDone: false,
	}
}
