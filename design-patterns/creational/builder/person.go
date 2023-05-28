package builder

type Person struct {
    name string
    age int
    height int
}

type PersonBuilder struct {
    person Person
}

func NewPersonBuilder() *PersonBuilder {
    return &PersonBuilder{
        person: Person{},
    }
}

func (p *PersonBuilder) name(name string) *PersonBuilder {
    p.person.name = name
    return p
}

func (p *PersonBuilder) age(age int) *PersonBuilder {
    p.person.age = age
    return p
}

func (p *PersonBuilder) height(height int) *PersonBuilder {
    p.person.height = height
    return p
}

func (p *PersonBuilder) build() Person {
    return p.person
}