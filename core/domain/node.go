package domain

type Node struct {
	Id   string
	Name string
}

type Relationship struct {
	NameA string
	NameB string
	Kind  string
	Color string
	Id    string
}

type Collection struct {
	Nodes         map[string]Node
	Relationships []Relationship
}
