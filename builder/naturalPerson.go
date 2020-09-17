package builder

//NaturalPerson struct
type NaturalPerson struct {
	Name     string
	Document string
	Type     string
}

func (person *NaturalPerson) SetName(name string) {
	person.Name = name
}

func (person *NaturalPerson) SetDocument(document string) {
	person.Document = document
}

func (person *NaturalPerson) Build() Person {
	return Person{
		Name:     person.Name,
		Document: person.Document,
		Type:     person.Type,
	}
}
