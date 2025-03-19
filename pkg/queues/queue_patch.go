package queues

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a request to update an existing queue.
type QueuePatch struct {
	// The display name. This may be used as a more human-readable name.
	DisplayName *util.Nullable[string] `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The description. This may be used as a space for notes or other information about the queue.
	Description *util.Nullable[string] `json:"description,omitempty" maxLength:"500" pattern:"^.*$"`
}

func (q *QueuePatch) GetDisplayName() *util.Nullable[string] {
	if q == nil {
		return nil
	}
	return q.DisplayName
}

func (q *QueuePatch) SetDisplayName(displayName util.Nullable[string]) {
	q.DisplayName = &displayName
}

func (q *QueuePatch) SetDisplayNameNull() {
	q.DisplayName = &util.Nullable[string]{IsNull: true}
}

func (q *QueuePatch) GetDescription() *util.Nullable[string] {
	if q == nil {
		return nil
	}
	return q.Description
}

func (q *QueuePatch) SetDescription(description util.Nullable[string]) {
	q.Description = &description
}

func (q *QueuePatch) SetDescriptionNull() {
	q.Description = &util.Nullable[string]{IsNull: true}
}

func (q QueuePatch) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueuePatch to string"
	}
	return string(jsonData)
}
