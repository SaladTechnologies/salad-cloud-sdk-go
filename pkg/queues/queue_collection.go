package queues

import "encoding/json"

// Represents a Queue Collection
type QueueCollection struct {
	// The list of queues.
	Items []Queue `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (q *QueueCollection) GetItems() []Queue {
	if q == nil {
		return nil
	}
	return q.Items
}

func (q *QueueCollection) SetItems(items []Queue) {
	q.Items = items
}

func (q QueueCollection) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueCollection to string"
	}
	return string(jsonData)
}
