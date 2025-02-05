package inferenceendpoints

import (
	"encoding/json"
)

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
	touched   map[string]bool
}

func (i *InferenceEndpoint) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InferenceEndpoint) SetId(id string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Id"] = true
	i.Id = &id
}

func (i *InferenceEndpoint) SetIdNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Id"] = true
	i.Id = nil
}

func (i *InferenceEndpoint) GetName() *string {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *InferenceEndpoint) SetName(name string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Name"] = true
	i.Name = &name
}

func (i *InferenceEndpoint) SetNameNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Name"] = true
	i.Name = nil
}

func (i *InferenceEndpoint) GetDisplayName() *string {
	if i == nil {
		return nil
	}
	return i.DisplayName
}

func (i *InferenceEndpoint) SetDisplayName(displayName string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["DisplayName"] = true
	i.DisplayName = &displayName
}

func (i *InferenceEndpoint) SetDisplayNameNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["DisplayName"] = true
	i.DisplayName = nil
}

func (i *InferenceEndpoint) GetDescription() *string {
	if i == nil {
		return nil
	}
	return i.Description
}

func (i *InferenceEndpoint) SetDescription(description string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Description"] = true
	i.Description = &description
}

func (i *InferenceEndpoint) SetDescriptionNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Description"] = true
	i.Description = nil
}

func (i *InferenceEndpoint) GetEndpointUrl() *string {
	if i == nil {
		return nil
	}
	return i.EndpointUrl
}

func (i *InferenceEndpoint) SetEndpointUrl(endpointUrl string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["EndpointUrl"] = true
	i.EndpointUrl = &endpointUrl
}

func (i *InferenceEndpoint) SetEndpointUrlNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["EndpointUrl"] = true
	i.EndpointUrl = nil
}

func (i *InferenceEndpoint) GetReadme() *string {
	if i == nil {
		return nil
	}
	return i.Readme
}

func (i *InferenceEndpoint) SetReadme(readme string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Readme"] = true
	i.Readme = &readme
}

func (i *InferenceEndpoint) SetReadmeNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Readme"] = true
	i.Readme = nil
}

func (i *InferenceEndpoint) GetPriceDescription() *string {
	if i == nil {
		return nil
	}
	return i.PriceDescription
}

func (i *InferenceEndpoint) SetPriceDescription(priceDescription string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["PriceDescription"] = true
	i.PriceDescription = &priceDescription
}

func (i *InferenceEndpoint) SetPriceDescriptionNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["PriceDescription"] = true
	i.PriceDescription = nil
}

func (i *InferenceEndpoint) GetIconImage() *string {
	if i == nil {
		return nil
	}
	return i.IconImage
}

func (i *InferenceEndpoint) SetIconImage(iconImage string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["IconImage"] = true
	i.IconImage = &iconImage
}

func (i *InferenceEndpoint) SetIconImageNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["IconImage"] = true
	i.IconImage = nil
}

func (i InferenceEndpoint) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["Id"] && i.Id == nil {
		data["id"] = nil
	} else if i.Id != nil {
		data["id"] = i.Id
	}

	if i.touched["Name"] && i.Name == nil {
		data["name"] = nil
	} else if i.Name != nil {
		data["name"] = i.Name
	}

	if i.touched["DisplayName"] && i.DisplayName == nil {
		data["display_name"] = nil
	} else if i.DisplayName != nil {
		data["display_name"] = i.DisplayName
	}

	if i.touched["Description"] && i.Description == nil {
		data["description"] = nil
	} else if i.Description != nil {
		data["description"] = i.Description
	}

	if i.touched["EndpointUrl"] && i.EndpointUrl == nil {
		data["endpoint_url"] = nil
	} else if i.EndpointUrl != nil {
		data["endpoint_url"] = i.EndpointUrl
	}

	if i.touched["Readme"] && i.Readme == nil {
		data["readme"] = nil
	} else if i.Readme != nil {
		data["readme"] = i.Readme
	}

	if i.touched["PriceDescription"] && i.PriceDescription == nil {
		data["price_description"] = nil
	} else if i.PriceDescription != nil {
		data["price_description"] = i.PriceDescription
	}

	if i.touched["IconImage"] && i.IconImage == nil {
		data["icon_image"] = nil
	} else if i.IconImage != nil {
		data["icon_image"] = i.IconImage
	}

	return json.Marshal(data)
}

func (i InferenceEndpoint) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpoint to string"
	}
	return string(jsonData)
}
