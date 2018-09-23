package ast

// File represents a file that is being compiled
type File struct {
	Name       string
	Statements []Statement
}

// Length returns the list of statements in the file
func (f *File) Length() int {
	// TODO: this will have to do something to recurse down the chain and figure out blocks and add that to the total
	// return len(f.Statements)

	// for _, stmt := range f.Statements {
	// 	// TODO: statement should define a .Length() function that will return the length of the statement node
	// }

	return -1
}

// NewFile retuns a new file and sets the filename
func NewFile(filename string) *File {
	return &File{
		Name: filename,
	}
}

// AddStatement appends a statement to the file
func (f *File) AddStatement(stmt Statement) {
	f.Statements = append(f.Statements, stmt)
}

func (f *File) Kind() NodeType { return FileNode }
