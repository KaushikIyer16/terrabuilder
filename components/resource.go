package components

import "fmt"

type Resource struct {
	Name     string
	Type     string
	Simple   []SimpleToken
	Compound []CompoundToken
}

func (resource *Resource) ToToken() string {
	return resource.ToResource()
}

func (resource *Resource) AddCompound(token *CompoundToken) *Resource {
	resource.Compound = append(resource.Compound, *token)
	return resource
}

func (resource *Resource) AddSimple(name, value string) *Resource {
	resource.Simple = append(resource.Simple, *NewSimpleToken(name, value))
	return resource
}

func (resource *Resource) ToResource() string {
	instance := fmt.Sprintf("resource \"%s\" \"%s\" {\n", resource.Type, resource.Name)
	/* loop through Simple Token*/
	for _, token := range resource.Simple {
		tokenString := token.ToToken()
		instance += fmt.Sprintf("%*s", len(tokenString)+2, tokenString)
	}
	/* loop through Compound Token*/
	for _, token := range resource.Compound {
		instance += token.ToToken(2)
	}
	instance += "}\n"
	return instance
}

func NewResource(name string, resourceType string) *Resource {
	return &Resource{Name:name, Type:resourceType, Simple: []SimpleToken{}, Compound: []CompoundToken{}}
}

