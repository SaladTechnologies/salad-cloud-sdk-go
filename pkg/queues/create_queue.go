package queues

// Represents a request to create a new queue.
type CreateQueue struct {
	// The queue name. This must be unique within the project.
	Name *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The display name. This may be used as a more human-readable name.
	DisplayName *string `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The description. This may be used as a space for notes or other information about the queue.
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
