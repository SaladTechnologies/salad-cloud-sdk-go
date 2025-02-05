package queues

import (
	"encoding/json"
)

// Represents a request to update an existing queue.
type UpdateQueue struct {
	// The display name. This may be used as a more human-readable name.
	DisplayName *string `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The description. This may be used as a space for notes or other information about the queue.
	Description *string `json:"description,omitempty" maxLength:"500"`
	touched     map[string]bool
}

func (u *UpdateQueue) GetDisplayName() *string {
	if u == nil {
		return nil
	}
	return u.DisplayName
}

func (u *UpdateQueue) SetDisplayName(displayName string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["DisplayName"] = true
	u.DisplayName = &displayName
}

func (u *UpdateQueue) SetDisplayNameNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["DisplayName"] = true
	u.DisplayName = nil
}

func (u *UpdateQueue) GetDescription() *string {
	if u == nil {
		return nil
	}
	return u.Description
}

func (u *UpdateQueue) SetDescription(description string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Description"] = true
	u.Description = &description
}

func (u *UpdateQueue) SetDescriptionNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Description"] = true
	u.Description = nil
}

func (u UpdateQueue) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["DisplayName"] && u.DisplayName == nil {
		data["display_name"] = nil
	} else if u.DisplayName != nil {
		data["display_name"] = u.DisplayName
	}

	if u.touched["Description"] && u.Description == nil {
		data["description"] = nil
	} else if u.Description != nil {
		data["description"] = u.Description
	}

	return json.Marshal(data)
}

func (u UpdateQueue) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateQueue to string"
	}
	return string(jsonData)
}
