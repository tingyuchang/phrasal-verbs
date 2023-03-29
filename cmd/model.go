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
