package template

import "testing"

func TestTemplateMethod(t *testing.T) {
	d := &AbstractDisplay{}
	d.Displayer = NewStringDisplay("hello world.")
	d.display()
	d.Displayer = NewCharDisplay('c')
	d.display()
}
