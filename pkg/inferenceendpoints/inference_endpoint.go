package inferenceendpoints

// Represents an inference endpoint
type InferenceEndpoint struct {
	// The unique identifier
	Id *string `json:"id,omitempty" required:"true"`
	// The inference endpoint name
	Name *string `json:"name,omitempty" required:"true"`
	// The inference endpoint display name
	DisplayName *string `json:"display_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// a brief description of the inference endpoint
	Description *string `json:"description,omitempty" required:"true"`
	// The URL of the inference endpoint
	EndpointUrl *string `json:"endpoint_url,omitempty" required:"true"`
	// A markdown file containing a detailed description of the inference endpoint
	Readme *string `json:"readme,omitempty" required:"true"`
	// A description of the price
	PriceDescription *string `json:"price_description,omitempty" required:"true"`
	// The URL of the icon image
	IconImage *string `json:"icon_image,omitempty" required:"true"`
}

func (i *InferenceEndpoint) SetId(id string) {
	i.Id = &id
}

func (i *InferenceEndpoint) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InferenceEndpoint) SetName(name string) {
	i.Name = &name
}

func (i *InferenceEndpoint) GetName() *string {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *InferenceEndpoint) SetDisplayName(displayName string) {
	i.DisplayName = &displayName
}

func (i *InferenceEndpoint) GetDisplayName() *string {
	if i == nil {
		return nil
	}
	return i.DisplayName
}

func (i *InferenceEndpoint) SetDescription(description string) {
	i.Description = &description
}

func (i *InferenceEndpoint) GetDescription() *string {
	if i == nil {
		return nil
	}
	return i.Description
}

func (i *InferenceEndpoint) SetEndpointUrl(endpointUrl string) {
	i.EndpointUrl = &endpointUrl
}

func (i *InferenceEndpoint) GetEndpointUrl() *string {
	if i == nil {
		return nil
	}
	return i.EndpointUrl
}

func (i *InferenceEndpoint) SetReadme(readme string) {
	i.Readme = &readme
}

func (i *InferenceEndpoint) GetReadme() *string {
	if i == nil {
		return nil
	}
	return i.Readme
}

func (i *InferenceEndpoint) SetPriceDescription(priceDescription string) {
	i.PriceDescription = &priceDescription
}

func (i *InferenceEndpoint) GetPriceDescription() *string {
	if i == nil {
		return nil
	}
	return i.PriceDescription
}

func (i *InferenceEndpoint) SetIconImage(iconImage string) {
	i.IconImage = &iconImage
}

func (i *InferenceEndpoint) GetIconImage() *string {
	if i == nil {
		return nil
	}
	return i.IconImage
}
