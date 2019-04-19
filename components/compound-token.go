package components

import "fmt"

type CompoundToken struct {
	Name     string
	Simple   []*SimpleToken
	Compound []*CompoundToken
}

func (token *CompoundToken) ToToken(parentTabs int) string {
	instance := fmt.Sprintf("%*s {\n", len(token.Name)+parentTabs, token.Name)
	/* loop through the Simple tokens*/
	for _, simpleToken := range token.Simple {
		tokenString := simpleToken.ToToken()
		instance += fmt.Sprintf("%*s", len(tokenString)+2+parentTabs, tokenString)
	}
	/* loop through the Compound tokens */
	for _, compoundToken := range token.Compound {
		instance += compoundToken.ToToken(parentTabs+2)
	}
	instance += fmt.Sprintf("%*s\n", parentTabs+1, "}")
	return instance
}

func (token *CompoundToken) Add(name, value string) *CompoundToken {
	token.Simple = append(token.Simple, NewSimpleToken(name, value))
	return token
}

func (token *CompoundToken) AddCompound(compoundToken *CompoundToken) *CompoundToken {
	 token.Compound = append(token.Compound, compoundToken)
	 return token
}

func NewCompoundToken(name string) *CompoundToken {
	return &CompoundToken{Name: name}
}
