package components

import "fmt"

type Resource struct {
	Name     string
	Type     string
	simple   []SimpleToken
	compound []CompoundToken
}

func (resource *Resource) ToToken() string {
	return resource.ToResource()
}

func (resource *Resource) AddCompound(token *CompoundToken) *Resource {
	resource.compound = append(resource.compound, *token)
	return resource
}

func (resource *Resource) AddSimple(name, value string) *Resource {
	resource.simple = append(resource.simple, NewSimpleToken(name, value))
	return resource
}

func (resource *Resource) ToResource() string {
	instance := fmt.Sprintf("resource \"%s\" \"%s\" {\n", resource.Type, resource.Name)
	/* loop through simple token*/
	for _, token := range resource.simple {
		tokenString := token.ToToken()
		instance += fmt.Sprintf("%*s", len(tokenString)+2, tokenString)
	}
	/* loop through compound token*/
	for _, token := range resource.compound {
		instance += token.ToToken(2)
	}
	instance += "}\n"
	return instance
}

func NewResource(name string, resourceType string) *Resource {
	return &Resource{Name:name, Type:resourceType, simple: []SimpleToken{}, compound: []CompoundToken{}}
}

