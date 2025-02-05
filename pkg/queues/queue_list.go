package queues

import (
	"encoding/json"
)

// Represents a list of queues
type QueueList struct {
	// The list of queues.
	Items   []Queue `json:"items,omitempty" required:"true" maxItems:"100"`
	touched map[string]bool
}

func (q *QueueList) GetItems() []Queue {
	if q == nil {
		return nil
	}
	return q.Items
}

func (q *QueueList) SetItems(items []Queue) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Items"] = true
	q.Items = items
}

func (q *QueueList) SetItemsNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Items"] = true
	q.Items = nil
}

func (q QueueList) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if q.touched["Items"] && q.Items == nil {
		data["items"] = nil
	} else if q.Items != nil {
		data["items"] = q.Items
	}

	return json.Marshal(data)
}

func (q QueueList) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueList to string"
	}
	return string(jsonData)
}
