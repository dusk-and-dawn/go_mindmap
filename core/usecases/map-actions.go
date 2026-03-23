package usecases

type MapActions struct {
	Prefix string
}

func (f MapActions) GetMapByID(id string) string {
	return f.Prefix + id
}
