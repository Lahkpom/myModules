package courses

type class struct {
	name        string
	description string
}

func NewClass(name string, description string) *class {
	return &class{
		name:        name,
		description: description,
	}
}
