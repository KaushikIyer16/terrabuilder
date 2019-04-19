package components

import (
	"fmt"
)

type Provider struct {
	name string
	token []SimpleToken
}

func (provider *Provider) ToToken() string {
	return provider.ToProvider()
}

func (provider *Provider) Add(name, value string) *Provider {
	provider.token = append(provider.token, NewSimpleToken(name, value))
	return provider
}

func (provider *Provider) AddVariable(name string,variable Variable) *Provider {
	provider.token = append(provider.token, NewSimpleToken(name, variable.ToReference()))
	return provider
}

func NewProvider(name string) *Provider {
	return &Provider{name:name, token: []SimpleToken{}}
}

func (provider *Provider) ToProvider() string{
	instance := fmt.Sprintf("provider \"%s\" {\n", provider.name)
	/* iterate through the map and insert them */
	for _, token := range provider.token  {
		tokenString := token.ToToken()
		instance += fmt.Sprintf("%*s", len(tokenString)+2, tokenString)
	}
	instance += "}\n"
	return instance
}