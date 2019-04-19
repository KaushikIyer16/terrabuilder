package components

import "fmt"

type Variable struct {
	Name string
	Value string
}

func (variable *Variable) ToToken() string {
	return fmt.Sprintf("var %s {}\n", variable.Name)
}

func (variable *Variable) ToVariable() string{
	return fmt.Sprintf("%s = %s\n", variable.Name, variable.Value)
}

func (variable *Variable) ToReference() string {
	return fmt.Sprintf("${var.%s}", variable.Name)
}