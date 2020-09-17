package proxy

import (
	"fmt"
)

//Application struct
type Application struct {
}

//Access a function that will do the access
func (a *Application) Access(url string, port string, user string, pass string) (string, error) {

	return fmt.Sprintf("Success to connect database in url: %s port: %s user: %s", url, port, user), nil
}
