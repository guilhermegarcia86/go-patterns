package proxy

import (
	"log"
	"time"
)

//Proxy struct
type Proxy struct {
	application *Application
	url         string
	port        string
	user        string
	pass        string
}

//OpenConnection a function that simulates a heavy call to create a database connection
func OpenConnection(url string, port string, user string, pass string) *Proxy {
	log.Println("A heavy process to create my connection with database")
	time.Sleep(2 * time.Second)
	return &Proxy{
		application: &Application{},
		url:         url,
		port:        port,
		user:        user,
		pass:        pass,
	}
}

//Access a function that receives params and verifies if has connection opened and do access
func (p *Proxy) Access(url string, port string, user string, pass string) (string, error) {

	if *p == (Proxy{}) {
		p = OpenConnection(url, port, user, pass)
	}

	msg, err := p.application.Access(url, port, user, pass)

	if err != nil {
		log.Fatalln("ERROR")
	}

	return msg, nil
}
