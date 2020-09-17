package chain

import (
	"log"

	"github.com/guilhermegarcia86/go-patterns/builder"
)

//ValidationName struct
type ValidationName struct {
	Next Validation
}

//Execute a function to implement the user name validation execution
func (validationName *ValidationName) Execute(user builder.Person) {

	if user.ValidationNameDone {
		log.Println("Validation name already done " + user.Name)
		validationName.Next.Execute(user)
		return
	}

	log.Println("Validation name user starting " + user.Name)
	user.ValidationNameDone = true
	log.Println("Validation finished")
}

//SetNext a function that set the next call
func (validationName *ValidationName) SetNext(next Validation) {
	validationName.Next = next
}
