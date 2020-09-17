package abstractfactory

//PessoaFisica struct
type PessoaFisica struct {
}

//MakeUser implementation of AbstractFactory
func (pessoa *PessoaFisica) MakeUser(name string, document string, tipo string) IUser {
	return &User{
		Name:               name,
		Document:           document,
		Tipo:               tipo,
		ValidationNameDone: false,
		ValidationTypeDone: false,
	}
}
