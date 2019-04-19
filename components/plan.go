package components

type Plan struct {
	Variables map[string]Variable
	Providers map[string]Provider
	Resources map[string]map[string]Resource
}

func (plan *Plan) AddProvider(provider *Provider) *Plan {
	plan.Providers[provider.Name] = *provider
	return plan
}


func (plan *Plan) AddVariable(variable Variable) *Plan {
	plan.Variables[variable.Name] = variable
	return plan
}

func (plan *Plan) AddResource(provider *Provider, resource *Resource) *Plan {
	if plan.Resources[provider.Name] == nil {
		plan.Resources[provider.Name] = map[string]Resource{}
	}
	plan.Resources[provider.Name][resource.Name] = *resource
	return plan
}

func (plan *Plan) DeleteResource(providerName, resourceName string) *Plan {
	delete(plan.Resources[providerName], resourceName)
	return plan
}

func (plan *Plan) ToToken() string {
	instance := ""
	/*loop through Variables*/
	for _, variable := range plan.Variables {
		instance += variable.ToToken()
	}
	/* loop through Providers*/
	for _, provider := range plan.Providers {
		instance += provider.ToToken()
	}
	/* loop through the map and list out the Resources */
	for _, resourceList := range plan.Resources {
		for _, resource := range resourceList {
			instance += resource.ToToken()
		}
	}
	return instance
}

func NewPlan() *Plan {
	return &Plan{Resources: map[string]map[string]Resource{}, Providers: map[string]Provider{}, Variables: map[string]Variable{}}
}