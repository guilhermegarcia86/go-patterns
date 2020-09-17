package chain

import (
	"fmt"

	abstractfactory "github.com/guilhermegarcia86/go-patterns/abstractFactory"
)

//ValidationType struct
type ValidationType struct {
	Next Validation
}

//Execute a function to implement the user type validation execution
func (validationType *ValidationType) Execute(user *abstractfactory.User) {

	if user.ValidationTypeDone {
		fmt.Println("Validation type already done")
		validationType.Next.Execute(user)
		return
	}

	fmt.Println("Validation type user starting")
	user.ValidationTypeDone = true
	validationType.Next.Execute(user)
}

//SetNext a function that set the next call
func (validationType *ValidationType) SetNext(next Validation) {
	validationType.Next = next
}
