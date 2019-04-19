package components

import "fmt"

type CompoundToken struct {
	name string
	simple   []SimpleToken
	compound []*CompoundToken
}

func (token *CompoundToken) ToToken(parentTabs int) string {
	instance := fmt.Sprintf("%*s {\n", len(token.name)+parentTabs, token.name)
	/* loop through the simple tokens*/
	for _, simpleToken := range token.simple {
		tokenString := simpleToken.ToToken()
		instance += fmt.Sprintf("%*s", len(tokenString)+2+parentTabs, tokenString)
	}
	/* loop through the compound tokens */
	for _, compoundToken := range token.compound {
		instance += compoundToken.ToToken(parentTabs+2)
	}
	instance += fmt.Sprintf("%*s\n", parentTabs+1, "}")
	return instance
}

func (token *CompoundToken) Add(name, value string) *CompoundToken {
	token.simple = append(token.simple, NewSimpleToken(name, value))
	return token
}

func (token *CompoundToken) AddCompound(compoundToken *CompoundToken) *CompoundToken {
	 token.compound = append(token.compound, compoundToken)
	 return token
}

func NewCompoundToken(name string) *CompoundToken {
	return &CompoundToken{name: name}
}
