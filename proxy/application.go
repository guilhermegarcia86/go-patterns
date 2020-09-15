package proxy

import (
	"fmt"
)

type Application struct {
}

func (a *Application) Access(url string, port string, user string, pass string) (string, error) {

	return fmt.Sprintf("Success to connect database in url: %s port: %s user: %s", url, port, user), nil
}
