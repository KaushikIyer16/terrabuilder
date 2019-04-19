package components

import "fmt"

type SimpleToken struct {
	name string
	value string
}

func (simpleToken *SimpleToken) ToToken() string {
	return fmt.Sprintf("%s = \"%s\"\n", simpleToken.name, simpleToken.value)
}

func NewSimpleToken(name, value string) SimpleToken {
	return SimpleToken{name:name, value:value}
}

