package components

import "fmt"

type SimpleToken struct {
	Name  string
	Value string
}

func (simpleToken *SimpleToken) ToToken() string {
	return fmt.Sprintf("%s = \"%s\"\n", simpleToken.Name, simpleToken.Value)
}

func NewSimpleToken(name, value string) *SimpleToken {
	return &SimpleToken{Name: name, Value:value}
}

