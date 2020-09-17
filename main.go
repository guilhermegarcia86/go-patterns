package main

import (
	"log"

	abstractfactory "github.com/guilhermegarcia86/go-patterns/abstractFactory"
	"github.com/guilhermegarcia86/go-patterns/chain"
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

	pfFactory, _ := abstractfactory.UserFactory("PF")
	pessoaFisica := pfFactory.MakeUser("Pessoa Fisica", "20411992023", "PF")

	pjFactory, _ := abstractfactory.UserFactory("PJ")
	pessoaJuridica := pjFactory.MakeUser("Pessoa Juridica", "82630818000152", "PJ")

	validationName := &chain.ValidationName{}
	validationType := &chain.ValidationType{}

	validationType.SetNext(validationName)
	validationType.Execute(pessoaFisica.GetUser())
	validationType.Execute(pessoaJuridica.GetUser())
}
