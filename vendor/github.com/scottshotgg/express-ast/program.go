package ast

// Program represents the following form:
// [ statement ]*
type Program struct {
	// Name string // Don't know if I should include this or not
	Files map[string]*File
}

// Length returns the length of files in the program
func (p *Program) Length() int {
	return len(p.Files)
}

func NewProgram() *Program {
	return &Program{
		Files: map[string]*File{},
	}
}

func (p *Program) AddFile(f *File) {
	p.Files[f.Name] = f
}
