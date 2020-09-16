package main

import (
	"fmt"
	"reflect"

	abstractfactory "github.com/guilhermegarcia86/go-patterns/abstractFactory"
)

func main() {

	// const (
	// 	url  = "urlDatabase"
	// 	port = "3306"
	// 	user = "user"
	// 	pass = "pass"
	// )

	// conn := proxy.OpenConnection(url, port, user, pass)

	// msgI, err := conn.Access(url, port, user, pass)
	// if err != nil {
	// 	log.Fatalln("ERROR")
	// }
	// log.Println(msgI)

	// msgII, err := conn.Access(url, port, user, pass)
	// if err != nil {
	// 	log.Fatal("ERROR")
	// }
	// log.Println(msgII)

	// pfFactory, _ := abstractfactory.UserFactory("PF")
	// pessoaFisica := pfFactory.MakeUser("Pessoa Fisica", "20411992023", "PF")
	// printUserDetails(pessoaFisica)

	// pjFactory, _ := abstractfactory.UserFactory("PJ")
	// pessoaJuridica := pjFactory.MakeUser("Pessoa Juridica", "82630818000152", "PJ")
	// printUserDetails(pessoaJuridica)

	// printType(pessoaFisica)
	// printType(pessoaJuridica)
}

func printType(t interface{}) {
	fmt.Println(reflect.TypeOf(t))
}

func printUserDetails(user abstractfactory.IUser) {
	fmt.Printf("Name: %s", user.GetName())
	fmt.Println()
	fmt.Printf("Document: %s", user.GetDocument())
	fmt.Println()
	fmt.Printf("Tipo: %s", user.GetTipo())
	fmt.Println()
}

// func main() {
// 	adidasFactory, _ := getSportsFactory("adidas")
// 	nikeFactory, _ := getSportsFactory("nike")
// 	nikeShort := nikeFactory.makeShort()
// 	adidasShort := adidasFactory.makeShort()
// 	printShortDetails(nikeShort)
// 	printShortDetails(adidasShort)
// }

// func printShortDetails(s iShort) {
// 	fmt.Printf("Logo: %s", s.getLogo())
// 	fmt.Println()
// 	fmt.Printf("Size: %d", s.getSize())
// 	fmt.Println()
// }

// type nikeShort struct {
// 	short
// }

// type adidasShort struct {
// 	short
// }

// type iShort interface {
// 	setLogo(logo string)
// 	setSize(size int)
// 	getLogo() string
// 	getSize() int
// }

// type short struct {
// 	logo string
// 	size int
// }

// func (s *short) setLogo(logo string) {
// 	s.logo = logo
// }

// func (s *short) getLogo() string {
// 	return s.logo
// }

// func (s *short) setSize(size int) {
// 	s.size = size
// }

// func (s *short) getSize() int {
// 	return s.size
// }

// type nike struct {
// }

// func (n *nike) makeShort() iShort {
// 	return &nikeShort{
// 		short: short{
// 			logo: "nike",
// 			size: 14,
// 		},
// 	}
// }

// type adidas struct {
// }

// func (a *adidas) makeShort() iShort {
// 	return &adidasShort{
// 		short: short{
// 			logo: "adidas",
// 			size: 14,
// 		},
// 	}
// }

// type iSportsFactory interface {
// 	makeShort() iShort
// }

// func getSportsFactory(brand string) (iSportsFactory, error) {
// 	if brand == "adidas" {
// 		return &adidas{}, nil
// 	}
// 	if brand == "nike" {
// 		return &nike{}, nil
// 	}
// 	return nil, fmt.Errorf("Wrong brand type passed")
// }
