package cmd

type PhrasalVerb struct {
	Name        string `csv:"name"`
	Description string `csv:"description"`
	Example1    string `csv:"ex1"`
	Example2    string `csv:"ex2"`
}

type PhrasalVerbs []PhrasalVerb

func (p PhrasalVerbs) Len() int {
	return len(p)
}

func (p PhrasalVerbs) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p PhrasalVerbs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func InitVerb(name, description, example1, example2 string) *PhrasalVerb {
	// Both of them should have a value
	if name == "" || description == "" {
		return nil
	}

	return &PhrasalVerb{
		Name:        name,
		Description: description,
		Example1:    example1,
		Example2:    example2,
	}
}
