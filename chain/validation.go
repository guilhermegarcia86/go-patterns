package chain

import (
	"github.com/guilhermegarcia86/go-patterns/builder"
)

//Validation interface has functions for Chain of Responsability Pattern
type Validation interface {
	Execute(builder.Person)
	SetNext(Validation)
}
