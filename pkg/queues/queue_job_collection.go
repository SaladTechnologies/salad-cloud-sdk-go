package queues

import "encoding/json"

// Represents a Queue Job Collection
type QueueJobCollection struct {
	// The list of queue jobs
	Items []QueueJob `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (q *QueueJobCollection) GetItems() []QueueJob {
	if q == nil {
		return nil
	}
	return q.Items
}

func (q *QueueJobCollection) SetItems(items []QueueJob) {
	q.Items = items
}

func (q QueueJobCollection) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueJobCollection to string"
	}
	return string(jsonData)
}
