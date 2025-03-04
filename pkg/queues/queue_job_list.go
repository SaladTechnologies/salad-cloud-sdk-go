package queues

import "encoding/json"

// Represents a list of queue jobs
type QueueJobList struct {
	Items []QueueJob `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (q *QueueJobList) GetItems() []QueueJob {
	if q == nil {
		return nil
	}
	return q.Items
}

func (q *QueueJobList) SetItems(items []QueueJob) {
	q.Items = items
}

func (q QueueJobList) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueJobList to string"
	}
	return string(jsonData)
}
