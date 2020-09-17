package abstractfactory

import "errors"

//AbstractFactory abstract factory
type AbstractFactory interface {
	MakeUser(name string, document string, tipo string) IUser
}

//UserFactory factory that decides how struct to create
func UserFactory(tipo string) (AbstractFactory, error) {
	if tipo == "PF" {
		return &PessoaFisica{}, nil
	}
	if tipo == "PJ" {
		return &PessoaJuridica{}, nil
	}
	return nil, errors.New("Wrong type passed")
}
