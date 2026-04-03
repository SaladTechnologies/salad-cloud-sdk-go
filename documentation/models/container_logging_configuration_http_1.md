# ContainerLoggingConfigurationHttp1

Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.

**Properties**

| Name        | Type                                                                            | Required | Description                                                    |
| :---------- | :------------------------------------------------------------------------------ | :------- | :------------------------------------------------------------- |
| Compression | [shared.ContainerLoggingHttpCompression](container_logging_http_compression.md) | ✅       | The compression algorithm to apply to logs before transmission |
| Format      | [shared.ContainerLoggingHttpFormat](container_logging_http_format.md)           | ✅       | The format in which logs will be delivered                     |
| Headers     | [][shared.ContainerLoggingHttpHeader](container_logging_http_header.md)         | ✅       | Optional HTTP headers to include in log transmission requests  |
| Host        | string                                                                          | ✅       | The hostname or IP address of the HTTP logging endpoint        |
| Port        | int64                                                                           | ✅       | The port number of the HTTP logging endpoint (1-65535)         |
| Password    | string                                                                          | ❌       | Optional password for HTTP authentication                      |
| Path        | string                                                                          | ❌       | Optional URL path for the HTTP endpoint                        |
| User        | string                                                                          | ❌       | Optional username for HTTP authentication                      |
