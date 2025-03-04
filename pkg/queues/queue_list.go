package queues

import "encoding/json"

// Represents a list of queues
type QueueList struct {
	// The list of queues.
	Items []Queue `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (q *QueueList) GetItems() []Queue {
	if q == nil {
		return nil
	}
	return q.Items
}

func (q *QueueList) SetItems(items []Queue) {
	q.Items = items
}

func (q QueueList) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueList to string"
	}
	return string(jsonData)
}
