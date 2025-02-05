package queues

import (
	"encoding/json"
)

// Represents a request to create a new queue.
type CreateQueue struct {
	// The queue name. This must be unique within the project.
	Name *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The display name. This may be used as a more human-readable name.
	DisplayName *string `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The description. This may be used as a space for notes or other information about the queue.
	Description *string `json:"description,omitempty" maxLength:"500"`
	touched     map[string]bool
}

func (c *CreateQueue) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateQueue) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *CreateQueue) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}

func (c *CreateQueue) GetDisplayName() *string {
	if c == nil {
		return nil
	}
	return c.DisplayName
}

func (c *CreateQueue) SetDisplayName(displayName string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DisplayName"] = true
	c.DisplayName = &displayName
}

func (c *CreateQueue) SetDisplayNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DisplayName"] = true
	c.DisplayName = nil
}

func (c *CreateQueue) GetDescription() *string {
	if c == nil {
		return nil
	}
	return c.Description
}

func (c *CreateQueue) SetDescription(description string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Description"] = true
	c.Description = &description
}

func (c *CreateQueue) SetDescriptionNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Description"] = true
	c.Description = nil
}

func (c CreateQueue) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Name"] && c.Name == nil {
		data["name"] = nil
	} else if c.Name != nil {
		data["name"] = c.Name
	}

	if c.touched["DisplayName"] && c.DisplayName == nil {
		data["display_name"] = nil
	} else if c.DisplayName != nil {
		data["display_name"] = c.DisplayName
	}

	if c.touched["Description"] && c.Description == nil {
		data["description"] = nil
	} else if c.Description != nil {
		data["description"] = c.Description
	}

	return json.Marshal(data)
}

func (c CreateQueue) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateQueue to string"
	}
	return string(jsonData)
}
