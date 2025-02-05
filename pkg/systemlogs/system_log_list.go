package systemlogs

import (
	"encoding/json"
)

// Represents a list of system logs
type SystemLogList struct {
	Items   []SystemLog `json:"items,omitempty" required:"true" maxItems:"50"`
	touched map[string]bool
}

func (s *SystemLogList) GetItems() []SystemLog {
	if s == nil {
		return nil
	}
	return s.Items
}

func (s *SystemLogList) SetItems(items []SystemLog) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Items"] = true
	s.Items = items
}

func (s *SystemLogList) SetItemsNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Items"] = true
	s.Items = nil
}

func (s SystemLogList) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Items"] && s.Items == nil {
		data["items"] = nil
	} else if s.Items != nil {
		data["items"] = s.Items
	}

	return json.Marshal(data)
}

func (s SystemLogList) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SystemLogList to string"
	}
	return string(jsonData)
}
