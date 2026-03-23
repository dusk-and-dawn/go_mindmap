package usecases

// Ist die Definition des Struct typen hier sinnvoll oder sollte ich diese auslagern?
type Node struct {
	Id       string
	Content  string
	Children []string
	X        float64
	Y        float64
}

type MapActions struct {
	// Prefix string
	Nodes map[string]Node
}

func (f MapActions) GetMapByID(id string) (Node, bool) {
	node, exists := f.Nodes[id]
	return node, exists
	// return f.Prefix + id
}
