package queues

import "encoding/json"

// Represents a request to create a new queue.
type QueuePrototype struct {
	// The queue name. This must be unique within the project.
	Name *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The display name. This may be used as a more human-readable name.
	DisplayName *string `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The description. This may be used as a space for notes or other information about the queue.
	Description *string `json:"description,omitempty" maxLength:"500" pattern:"^.*$"`
}

func (q *QueuePrototype) GetName() *string {
	if q == nil {
		return nil
	}
	return q.Name
}

func (q *QueuePrototype) SetName(name string) {
	q.Name = &name
}

func (q *QueuePrototype) GetDisplayName() *string {
	if q == nil {
		return nil
	}
	return q.DisplayName
}

func (q *QueuePrototype) SetDisplayName(displayName string) {
	q.DisplayName = &displayName
}

func (q *QueuePrototype) GetDescription() *string {
	if q == nil {
		return nil
	}
	return q.Description
}

func (q *QueuePrototype) SetDescription(description string) {
	q.Description = &description
}

func (q QueuePrototype) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueuePrototype to string"
	}
	return string(jsonData)
}
