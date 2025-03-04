package queues

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a request to update an existing queue.
type UpdateQueue struct {
	// The display name. This may be used as a more human-readable name.
	DisplayName *util.Nullable[string] `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The description. This may be used as a space for notes or other information about the queue.
	Description *util.Nullable[string] `json:"description,omitempty" maxLength:"500"`
}

func (u *UpdateQueue) GetDisplayName() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.DisplayName
}

func (u *UpdateQueue) SetDisplayName(displayName util.Nullable[string]) {
	u.DisplayName = &displayName
}

func (u *UpdateQueue) SetDisplayNameNull() {
	u.DisplayName = &util.Nullable[string]{IsNull: true}
}

func (u *UpdateQueue) GetDescription() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.Description
}

func (u *UpdateQueue) SetDescription(description util.Nullable[string]) {
	u.Description = &description
}

func (u *UpdateQueue) SetDescriptionNull() {
	u.Description = &util.Nullable[string]{IsNull: true}
}

func (u UpdateQueue) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateQueue to string"
	}
	return string(jsonData)
}
