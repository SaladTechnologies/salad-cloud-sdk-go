package systemlogs

import "encoding/json"

// Represents a list of system logs
type SystemLogList struct {
	// A list of system logs
	Items []SystemLog `json:"items,omitempty" required:"true" maxItems:"50"`
}

func (s *SystemLogList) GetItems() []SystemLog {
	if s == nil {
		return nil
	}
	return s.Items
}

func (s *SystemLogList) SetItems(items []SystemLog) {
	s.Items = items
}

func (s SystemLogList) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SystemLogList to string"
	}
	return string(jsonData)
}
