# DatadogLoggingConfiguration

Configuration for forwarding container logs to Datadog monitoring service.

**Properties**

| Name   | Type                                                                           | Required | Description                                                                           |
| :----- | :----------------------------------------------------------------------------- | :------- | :------------------------------------------------------------------------------------ |
| ApiKey | string                                                                         | ✅       | The Datadog API key used for authentication when sending logs.                        |
| Host   | string                                                                         | ✅       | The Datadog intake server host URL where logs will be sent.                           |
| Tags   | [][shared.DatadogTagForContainerLogging](datadog_tag_for_container_logging.md) | ✅       | Optional metadata tags to attach to logs for filtering and categorization in Datadog. |
