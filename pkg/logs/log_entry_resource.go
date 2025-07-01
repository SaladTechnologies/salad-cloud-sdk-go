package logs

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
)

// The resource associated with the log entry
type LogEntryResource struct {
	// The labels associated with the resource
	Labels map[string]*string `json:"labels,omitempty" required:"true"`
	// The type of the resource
	Type_ *string `json:"type,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LogEntryResource) GetLabels() map[string]*string {
	if l == nil {
		return nil
	}
	return l.Labels
}

func (l *LogEntryResource) SetLabels(labels map[string]*string) {
	l.Labels = utils.CloneMap(labels)
}

func (l *LogEntryResource) GetType_() *string {
	if l == nil {
		return nil
	}
	return l.Type_
}

func (l *LogEntryResource) SetType_(type_ string) {
	l.Type_ = &type_
}

func (l LogEntryResource) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LogEntryResource to string"
	}
	return string(jsonData)
}
