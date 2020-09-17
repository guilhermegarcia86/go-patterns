package chain

import abstractfactory "github.com/guilhermegarcia86/go-patterns/abstractFactory"

//Validation interface has functions for Chain of Responsability Pattern
type Validation interface {
	Execute(*abstractfactory.User)
	SetNext(Validation)
}
