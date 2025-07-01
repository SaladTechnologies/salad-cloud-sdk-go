package logs

import "encoding/json"

type LogEntry struct {
	// The log message in JSON format.
	JsonLog any `json:"json_log,omitempty"`
	// The parent span ID of the log entry
	ParentSpanId *string `json:"parent_span_id,omitempty" maxLength:"1000" minLength:"1"`
	// The time when the log entry was received
	ReceiveTime *string `json:"receive_time,omitempty" required:"true"`
	// The resource associated with the log entry
	Resource *LogEntryResource `json:"resource,omitempty" required:"true"`
	// The severity level of the log entry
	Severity *LogEntrySeverity `json:"severity,omitempty" required:"true"`
	// The span ID of the log entry
	SpanId *string `json:"span_Id,omitempty" maxLength:"1000" minLength:"1"`
	// The log message in text format.
	TextLog *string `json:"text_log,omitempty" maxLength:"10000"`
	// The timestamp of the log entry
	Time *string `json:"time,omitempty" required:"true"`
	// The trace ID of the log entry
	TraceId *string `json:"trace_Id,omitempty" maxLength:"1000" minLength:"1"`
}

func (l *LogEntry) GetJsonLog() any {
	if l == nil {
		return nil
	}
	return l.JsonLog
}

func (l *LogEntry) SetJsonLog(jsonLog any) {
	l.JsonLog = &jsonLog
}

func (l *LogEntry) GetParentSpanId() *string {
	if l == nil {
		return nil
	}
	return l.ParentSpanId
}

func (l *LogEntry) SetParentSpanId(parentSpanId string) {
	l.ParentSpanId = &parentSpanId
}

func (l *LogEntry) GetReceiveTime() *string {
	if l == nil {
		return nil
	}
	return l.ReceiveTime
}

func (l *LogEntry) SetReceiveTime(receiveTime string) {
	l.ReceiveTime = &receiveTime
}

func (l *LogEntry) GetResource() *LogEntryResource {
	if l == nil {
		return nil
	}
	return l.Resource
}

func (l *LogEntry) SetResource(resource LogEntryResource) {
	l.Resource = &resource
}

func (l *LogEntry) GetSeverity() *LogEntrySeverity {
	if l == nil {
		return nil
	}
	return l.Severity
}

func (l *LogEntry) SetSeverity(severity LogEntrySeverity) {
	l.Severity = &severity
}

func (l *LogEntry) GetSpanId() *string {
	if l == nil {
		return nil
	}
	return l.SpanId
}

func (l *LogEntry) SetSpanId(spanId string) {
	l.SpanId = &spanId
}

func (l *LogEntry) GetTextLog() *string {
	if l == nil {
		return nil
	}
	return l.TextLog
}

func (l *LogEntry) SetTextLog(textLog string) {
	l.TextLog = &textLog
}

func (l *LogEntry) GetTime() *string {
	if l == nil {
		return nil
	}
	return l.Time
}

func (l *LogEntry) SetTime(time string) {
	l.Time = &time
}

func (l *LogEntry) GetTraceId() *string {
	if l == nil {
		return nil
	}
	return l.TraceId
}

func (l *LogEntry) SetTraceId(traceId string) {
	l.TraceId = &traceId
}

func (l LogEntry) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LogEntry to string"
	}
	return string(jsonData)
}
