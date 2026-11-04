package domain

type Node struct {
	Id   string
	Name string
}

type Relationship struct {
	IdA   string
	IdB   string
	Kind  string
	Color string
	Id    string
}

type Collection struct {
	Id            string
	Name          string
	Nodes         map[string]Node
	Relationships []Relationship
}
