package queues

// Represents a request to update an existing queue.
type UpdateQueue struct {
	// The display name. This may be used as a more human-readable name.
	DisplayName *string `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The description. This may be used as a space for notes or other information about the queue.
	Description *string `json:"description,omitempty" maxLength:"500"`
}

func (u *UpdateQueue) SetDisplayName(displayName string) {
	u.DisplayName = &displayName
}

func (u *UpdateQueue) GetDisplayName() *string {
	if u == nil {
		return nil
	}
	return u.DisplayName
}

func (u *UpdateQueue) SetDescription(description string) {
	u.Description = &description
}

func (u *UpdateQueue) GetDescription() *string {
	if u == nil {
		return nil
	}
	return u.Description
}
