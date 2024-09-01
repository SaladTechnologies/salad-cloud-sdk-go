package queues

// Represents a list of queues
type QueueList struct {
	Items []Queue `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (q *QueueList) SetItems(items []Queue) {
	q.Items = items
}

func (q *QueueList) GetItems() []Queue {
	if q == nil {
		return nil
	}
	return q.Items
}
