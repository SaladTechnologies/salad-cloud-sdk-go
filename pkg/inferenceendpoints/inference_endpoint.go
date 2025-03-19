package inferenceendpoints

import "encoding/json"

// Represents an inference endpoint
type InferenceEndpoint struct {
	// The inference endpoint identifier.
	Id *string `json:"id,omitempty" required:"true"`
	// The inference endpoint name.
	Name *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The organization name.
	OrganizationName *string `json:"organization_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The display-friendly name of the resource.
	DisplayName *string `json:"display_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The detailed description of the resource.
	Description *string `json:"description,omitempty" required:"true" maxLength:"1000" pattern:"^.*$"`
	// A markdown file containing a detailed description of the inference endpoint
	Readme *string `json:"readme,omitempty" required:"true" maxLength:"100000" minLength:"1" pattern:"^.*$"`
	// A description of the price
	PriceDescription *string `json:"price_description,omitempty" required:"true" maxLength:"100" minLength:"1" pattern:"^.*$"`
	// The URL of the icon image
	IconUrl *string `json:"icon_url,omitempty" required:"true" maxLength:"2048" minLength:"1" pattern:"^.*$"`
	// The input schema
	InputSchema *string `json:"input_schema,omitempty" required:"true" maxLength:"100000" minLength:"1" pattern:"^.*$"`
	// The output schema
	OutputSchema *string `json:"output_schema,omitempty" required:"true" maxLength:"100000" minLength:"1" pattern:"^.*$"`
}

func (i *InferenceEndpoint) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InferenceEndpoint) SetId(id string) {
	i.Id = &id
}

func (i *InferenceEndpoint) GetName() *string {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *InferenceEndpoint) SetName(name string) {
	i.Name = &name
}

func (i *InferenceEndpoint) GetOrganizationName() *string {
	if i == nil {
		return nil
	}
	return i.OrganizationName
}

func (i *InferenceEndpoint) SetOrganizationName(organizationName string) {
	i.OrganizationName = &organizationName
}

func (i *InferenceEndpoint) GetDisplayName() *string {
	if i == nil {
		return nil
	}
	return i.DisplayName
}

func (i *InferenceEndpoint) SetDisplayName(displayName string) {
	i.DisplayName = &displayName
}

func (i *InferenceEndpoint) GetDescription() *string {
	if i == nil {
		return nil
	}
	return i.Description
}

func (i *InferenceEndpoint) SetDescription(description string) {
	i.Description = &description
}

func (i *InferenceEndpoint) GetReadme() *string {
	if i == nil {
		return nil
	}
	return i.Readme
}

func (i *InferenceEndpoint) SetReadme(readme string) {
	i.Readme = &readme
}

func (i *InferenceEndpoint) GetPriceDescription() *string {
	if i == nil {
		return nil
	}
	return i.PriceDescription
}

func (i *InferenceEndpoint) SetPriceDescription(priceDescription string) {
	i.PriceDescription = &priceDescription
}

func (i *InferenceEndpoint) GetIconUrl() *string {
	if i == nil {
		return nil
	}
	return i.IconUrl
}

func (i *InferenceEndpoint) SetIconUrl(iconUrl string) {
	i.IconUrl = &iconUrl
}

func (i *InferenceEndpoint) GetInputSchema() *string {
	if i == nil {
		return nil
	}
	return i.InputSchema
}

func (i *InferenceEndpoint) SetInputSchema(inputSchema string) {
	i.InputSchema = &inputSchema
}

func (i *InferenceEndpoint) GetOutputSchema() *string {
	if i == nil {
		return nil
	}
	return i.OutputSchema
}

func (i *InferenceEndpoint) SetOutputSchema(outputSchema string) {
	i.OutputSchema = &outputSchema
}

func (i InferenceEndpoint) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpoint to string"
	}
	return string(jsonData)
}
