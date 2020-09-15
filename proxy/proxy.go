package proxy

import (
	"log"
	"time"
)

type Proxy struct {
	application *Application
	url         string
	port        string
	user        string
	pass        string
}

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
