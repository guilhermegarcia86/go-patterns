package abstractfactory

//PessoaJuridica struct
type PessoaJuridica struct {
}

//MakeUser func
func (pessoa *PessoaJuridica) MakeUser(name string, document string, tipo string) IUser {
	return &PessoaJuridicaUser{
		User: User{
			name:     name,
			document: document,
			tipo:     tipo,
		},
	}
}
