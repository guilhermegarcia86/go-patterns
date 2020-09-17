package chain

import (
	"fmt"

	abstractfactory "github.com/guilhermegarcia86/go-patterns/abstractFactory"
)

//ValidationName struct
type ValidationName struct {
	Next Validation
}

//Execute a function to implement the user name validation execution
func (validationName *ValidationName) Execute(user *abstractfactory.User) {

	if user.ValidationNameDone {
		fmt.Println("Validation name already done")
		validationName.Next.Execute(user)
		return
	}

	fmt.Println("Validation name user starting")
	user.ValidationNameDone = true
	fmt.Println("Validation finished")
}

//SetNext a function that set the next call
func (validationName *ValidationName) SetNext(next Validation) {
	validationName.Next = next
}
