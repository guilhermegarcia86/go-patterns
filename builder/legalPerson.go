package builder

//NaturalPerson struct
type LegalPerson struct {
	Name     string
	Document string
	Type     string
}

func (person *LegalPerson) SetName(name string) {
	person.Name = name
}

func (person *LegalPerson) SetDocument(document string) {
	person.Document = document
}

func (person *LegalPerson) Build() Person {
	return Person{
		Name:     person.Name,
		Document: person.Document,
		Type:     person.Type,
	}
}
