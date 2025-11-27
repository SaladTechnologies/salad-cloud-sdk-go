# LogEntry

**Properties**

| Name         | Type                                           | Required | Description                                |
| :----------- | :--------------------------------------------- | :------- | :----------------------------------------- |
| ReceiveTime  | string                                         | ✅       | The time when the log entry was received   |
| Resource     | [logs.LogEntryResource](log_entry_resource.md) | ✅       | The resource associated with the log entry |
| Severity     | [logs.LogEntrySeverity](log_entry_severity.md) | ✅       | The severity level of the log entry        |
| Time         | string                                         | ✅       | The timestamp of the log entry             |
| JsonLog      | any                                            | ❌       | The log message in JSON format.            |
| ParentSpanId | string                                         | ❌       | The parent span ID of the log entry        |
| SpanId       | string                                         | ❌       | The span ID of the log entry               |
| TextLog      | string                                         | ❌       | The log message in text format.            |
| TraceId      | string                                         | ❌       | The trace ID of the log entry              |
