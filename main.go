package main

import (
	"log"

	"github.com/guilhermegarcia86/go-patterns/proxy"
)

func main() {

	const (
		url  = "urlDatabase"
		port = "3306"
		user = "user"
		pass = "pass"
	)

	conn := proxy.OpenConnection(url, port, user, pass)

	msgI, err := conn.Access(url, port, user, pass)
	if err != nil {
		log.Fatalln("ERROR")
	}
	log.Println(msgI)

	msgII, err := conn.Access(url, port, user, pass)
	if err != nil {
		log.Fatal("ERROR")
	}
	log.Println(msgII)

}
