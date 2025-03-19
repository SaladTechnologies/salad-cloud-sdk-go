# ContainerLoggingConfiguration

Configuration options for directing container logs to a logging provider. This schema enables you to specify a single logging destination for container output, supporting monitoring, debugging, and analytics use cases. Each provider has its own configuration parameters defined in the referenced schemas. Only one logging provider can be selected at a time.

**Properties**

| Name     | Type                                       | Required | Description                                                                                                                                                                               |
| :------- | :----------------------------------------- | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Axiom    | shared.AxiomLoggingConfiguration           | ❌       | Configuration settings for integrating container logs with the Axiom logging service. When specified, container logs will be forwarded to the Axiom instance defined by these parameters. |
| Datadog  | shared.DatadogLoggingConfiguration         | ❌       | Configuration for forwarding container logs to Datadog monitoring service.                                                                                                                |
| Http     | shared.ContainerHttpLoggingConfiguration   | ❌       | Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.                                                                |
| NewRelic | shared.NewRelicLoggingConfiguration        | ❌       | Configuration for sending container logs to New Relic's log management platform.                                                                                                          |
| Splunk   | shared.ContainerLoggingSplunkConfiguration | ❌       | Configuration settings for forwarding container logs to a Splunk instance.                                                                                                                |
| Tcp      | shared.TcpLoggingConfiguration             | ❌       | Configuration for forwarding container logs to a remote TCP endpoint                                                                                                                      |
