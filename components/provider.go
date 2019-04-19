package components

import (
	"fmt"
)

type Provider struct {
	Name  string
	Token []*SimpleToken
}

func (provider *Provider) ToToken() string {
	return provider.ToProvider()
}

func (provider *Provider) Add(name, value string) *Provider {
	provider.Token = append(provider.Token, NewSimpleToken(name, value))
	return provider
}

func (provider *Provider) AddVariable(name string,variable Variable) *Provider {
	provider.Token = append(provider.Token, NewSimpleToken(name, variable.ToReference()))
	return provider
}

func NewProvider(name string) *Provider {
	return &Provider{Name: name, Token: []*SimpleToken{}}
}

func (provider *Provider) ToProvider() string{
	instance := fmt.Sprintf("provider \"%s\" {\n", provider.Name)
	/* iterate through the map and insert them */
	for _, token := range provider.Token {
		tokenString := token.ToToken()
		instance += fmt.Sprintf("%*s", len(tokenString)+2, tokenString)
	}
	instance += "}\n"
	return instance
}