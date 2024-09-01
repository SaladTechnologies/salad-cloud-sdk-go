package queues

// Represents a request to create a queue
type CreateQueue struct {
	Name        *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	DisplayName *string `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The description
	Description *string `json:"description,omitempty" maxLength:"500"`
}

func (c *CreateQueue) SetName(name string) {
	c.Name = &name
}

func (c *CreateQueue) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateQueue) SetDisplayName(displayName string) {
	c.DisplayName = &displayName
}

func (c *CreateQueue) GetDisplayName() *string {
	if c == nil {
		return nil
	}
	return c.DisplayName
}

func (c *CreateQueue) SetDescription(description string) {
	c.Description = &description
}

func (c *CreateQueue) GetDescription() *string {
	if c == nil {
		return nil
	}
	return c.Description
}
