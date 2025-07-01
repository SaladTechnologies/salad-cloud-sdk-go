package logs

// The severity level of the log entry
type LogEntrySeverity string

const (
	LOG_ENTRY_SEVERITY_DEBUG     LogEntrySeverity = "debug"
	LOG_ENTRY_SEVERITY_INFO      LogEntrySeverity = "info"
	LOG_ENTRY_SEVERITY_NOTICE    LogEntrySeverity = "notice"
	LOG_ENTRY_SEVERITY_WARNING   LogEntrySeverity = "warning"
	LOG_ENTRY_SEVERITY_ERROR     LogEntrySeverity = "error"
	LOG_ENTRY_SEVERITY_CRITICAL  LogEntrySeverity = "critical"
	LOG_ENTRY_SEVERITY_ALERT     LogEntrySeverity = "alert"
	LOG_ENTRY_SEVERITY_EMERGENCY LogEntrySeverity = "emergency"
)
