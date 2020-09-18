package main

import (
	"log"

	"github.com/guilhermegarcia86/go-patterns/builder"
	"github.com/guilhermegarcia86/go-patterns/chain"
	"github.com/guilhermegarcia86/go-patterns/proxy"
)

func main() {

	//Build a PersonBuilder, set values and build a NaturalPerson
	personBuilder := builder.GetBuilder("PF")

	personBuilder.SetName("John Doe")
	personBuilder.SetDocument("21368063004")
	naturalPerson := personBuilder.Build()

	//Build a PersonBuilder, set values and build a LegalPerson
	personBuilder = builder.GetBuilder("PJ")
	personBuilder.SetName("Cool Company")
	personBuilder.SetDocument("47902850000149")
	legalPerson := personBuilder.Build()

	//Begin a Validation chain
	validationName := &chain.ValidationName{}
	validationType := &chain.ValidationType{}

	validationType.SetNext(validationName)

	validationType.Execute(naturalPerson)

	validationType.Execute(legalPerson)

	const (
		url  = "urlDatabase"
		port = "3306"
		user = "user"
		pass = "pass"
	)

	//Open connection, heavy process
	conn := proxy.OpenConnection(url, port, user, pass)

	//Access database
	msgI, err := conn.Access(url, port, user, pass)
	if err != nil {
		log.Fatalln("ERROR ", err)
	}
	log.Println(msgI)

	//Do not open connection again
	msgII, err := conn.Access(url, port, user, pass)
	if err != nil {
		log.Fatal("ERROR ", err)
	}
	log.Println(msgII)

}
