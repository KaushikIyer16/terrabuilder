package components

type Plan struct {
	variables map[string]Variable
	providers map[string]Provider
	resources map[string][]Resource
}

func (plan *Plan) AddProvider(provider *Provider) *Plan {
	plan.providers[provider.name] = *provider
	return plan
}


func (plan *Plan) AddVariable(variable Variable) *Plan {
	plan.variables[variable.Name] = variable
	return plan
}

func (plan *Plan) AddResource(provider *Provider, resource *Resource) *Plan {
	plan.resources[provider.name] = append(plan.resources[provider.name], *resource)
	return plan
}

func (plan *Plan) ToToken() string {
	instance := ""
	/*loop through variables*/
	for _, variable := range plan.variables {
		instance += variable.ToToken()
	}
	/* loop through providers*/
	for _, provider := range plan.providers {
		instance += provider.ToToken()
	}
	/* loop through the map and list out the resources */
	for _, resourceList := range plan.resources {
		for _, resource := range resourceList {
			instance += resource.ToToken()
		}
	}
	return instance
}

func NewPlan() *Plan {
	return &Plan{resources: map[string][]Resource{}, providers: map[string]Provider{}, variables: map[string]Variable{}}
}