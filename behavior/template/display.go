package template

type Displayer interface {
	open()
	print()
	close()
}

type AbstractDisplay struct {
	Displayer Displayer
}

func (a *AbstractDisplay) display() {
	a.Displayer.open()
	for i := 0; i < 2; i++ {
		a.Displayer.print()
	}
	a.Displayer.close()
}

type StringDisplay struct {
	width int
	str   string
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		width: len(str),
		str:   str,
	}
}

func (s *StringDisplay) open() {
	s.printLine()
}

func (s *StringDisplay) print() {
	println("|" + s.str + "|")
}

func (s *StringDisplay) close() {
	s.printLine()
}

func (s *StringDisplay) printLine() {
	print("+")
	for i := 0; i < s.width; i++ {
		print("-")
	}
	println("+")
}

type CharDisplay struct {
	char byte
}

func NewCharDisplay(char byte) *CharDisplay {
	return &CharDisplay{
		char: char,
	}
}

func (c *CharDisplay) open() {
	print("<<")
}

func (c *CharDisplay) print() {
	print(string(c.char))
}

func (c *CharDisplay) close() {
	println("<<")
}
