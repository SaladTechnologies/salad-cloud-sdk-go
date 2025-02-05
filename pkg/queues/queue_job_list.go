package queues

import (
	"encoding/json"
)

// Represents a list of queue jobs
type QueueJobList struct {
	Items   []QueueJob `json:"items,omitempty" required:"true" maxItems:"100"`
	touched map[string]bool
}

func (q *QueueJobList) GetItems() []QueueJob {
	if q == nil {
		return nil
	}
	return q.Items
}

func (q *QueueJobList) SetItems(items []QueueJob) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Items"] = true
	q.Items = items
}

func (q *QueueJobList) SetItemsNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Items"] = true
	q.Items = nil
}

func (q QueueJobList) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if q.touched["Items"] && q.Items == nil {
		data["items"] = nil
	} else if q.Items != nil {
		data["items"] = q.Items
	}

	return json.Marshal(data)
}

func (q QueueJobList) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueJobList to string"
	}
	return string(jsonData)
}
