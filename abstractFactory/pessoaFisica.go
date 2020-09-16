package abstractfactory

//PessoaFisica struct
type PessoaFisica struct {
}

//MakeUser func
func (pessoa *PessoaFisica) MakeUser(name string, document string, tipo string) IUser {
	return &PessoaFisicaUser{
		User: User{
			name:     name,
			document: document,
			tipo:     tipo,
		},
	}
}
