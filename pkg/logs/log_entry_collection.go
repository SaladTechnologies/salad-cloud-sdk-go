package logs

import "encoding/json"

// Represents a page of organization logs
type LogEntryCollection struct {
	// A collection of log entries
	Items []LogEntry `json:"items,omitempty" required:"true" maxItems:"10000"`
	// The organization name.
	OrganizationName *string `json:"organization_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The maximum time page boundary. This may be used when getting paginated results.
	PageMaxTime *string `json:"page_max_time,omitempty" required:"true"`
	// The minimum time page boundary. This may be used when getting paginated results.
	PageMinTime *string `json:"page_min_time,omitempty" required:"true"`
}

func (l *LogEntryCollection) GetItems() []LogEntry {
	if l == nil {
		return nil
	}
	return l.Items
}

func (l *LogEntryCollection) SetItems(items []LogEntry) {
	l.Items = items
}

func (l *LogEntryCollection) GetOrganizationName() *string {
	if l == nil {
		return nil
	}
	return l.OrganizationName
}

func (l *LogEntryCollection) SetOrganizationName(organizationName string) {
	l.OrganizationName = &organizationName
}

func (l *LogEntryCollection) GetPageMaxTime() *string {
	if l == nil {
		return nil
	}
	return l.PageMaxTime
}

func (l *LogEntryCollection) SetPageMaxTime(pageMaxTime string) {
	l.PageMaxTime = &pageMaxTime
}

func (l *LogEntryCollection) GetPageMinTime() *string {
	if l == nil {
		return nil
	}
	return l.PageMinTime
}

func (l *LogEntryCollection) SetPageMinTime(pageMinTime string) {
	l.PageMinTime = &pageMinTime
}

func (l LogEntryCollection) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LogEntryCollection to string"
	}
	return string(jsonData)
}
