package chain

import (
	"log"

	"github.com/guilhermegarcia86/go-patterns/builder"
)

//ValidationType struct
type ValidationType struct {
	Next Validation
}

//Execute a function to implement the user type validation execution
func (validationType *ValidationType) Execute(user builder.Person) {

	if user.ValidationTypeDone {
		log.Println("Validation type already done " + user.Type)
		validationType.Next.Execute(user)
		return
	}

	log.Println("Validation type user starting " + user.Type)
	user.ValidationTypeDone = true
	log.Println("Execute next validation")
	validationType.Next.Execute(user)
}

//SetNext a function that set the next call
func (validationType *ValidationType) SetNext(next Validation) {
	validationType.Next = next
}
