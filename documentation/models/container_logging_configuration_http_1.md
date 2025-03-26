# ContainerLoggingConfigurationHttp1

Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.

**Properties**

| Name        | Type                                                 | Required | Description                                                    |
| :---------- | :--------------------------------------------------- | :------- | :------------------------------------------------------------- |
| Host        | string                                               | ✅       | The hostname or IP address of the HTTP logging endpoint        |
| Port        | int64                                                | ✅       | The port number of the HTTP logging endpoint (1-65535)         |
| Format      | shared.ContainerHttpLoggingConfigurationFormat1      | ✅       | The format in which logs will be delivered                     |
| Headers     | []shared.ContainerLoggingHttpHeader                  | ✅       | Optional HTTP headers to include in log transmission requests  |
| Compression | shared.ContainerHttpLoggingConfigurationCompression1 | ✅       | The compression algorithm to apply to logs before transmission |
| User        | string                                               | ❌       | Optional username for HTTP authentication                      |
| Password    | string                                               | ❌       | Optional password for HTTP authentication                      |
| Path        | string                                               | ❌       | Optional URL path for the HTTP endpoint                        |

# ContainerHttpLoggingConfigurationFormat1

The format in which logs will be delivered

**Properties**

| Name       | Type   | Required | Description  |
| :--------- | :----- | :------- | :----------- |
| json       | string | ✅       | "json"       |
| json_lines | string | ✅       | "json_lines" |

# ContainerHttpLoggingConfigurationCompression1

The compression algorithm to apply to logs before transmission

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| none | string | ✅       | "none"      |
| gzip | string | ✅       | "gzip"      |
