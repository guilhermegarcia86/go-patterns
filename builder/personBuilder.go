package builder

//PersonBuilder bulder
type PersonBuilder interface {
	SetName(name string)
	SetDocument(document string)
	Build() Person
}

//GetBuilder init builder
func GetBuilder(tyeName string) PersonBuilder {
	if tyeName == "PF" {
		return &NaturalPerson{
			Type: "PF",
		}
	}
	if tyeName == "PJ" {
		return &LegalPerson{
			Type: "PJ",
		}
	}
	return nil
}
