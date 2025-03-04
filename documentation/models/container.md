# Container

Represents a container

**Properties**

| Name                 | Type                                 | Required | Description                                  |
| :------------------- | :----------------------------------- | :------- | :------------------------------------------- |
| Image                | string                               | ✅       |                                              |
| Resources            | shared.ContainerResourceRequirements | ✅       | Represents a container resource requirements |
| Command              | []string                             | ✅       |                                              |
| Priority             | shared.ContainerGroupPriority        | ❌       |                                              |
| Size                 | int64                                | ❌       |                                              |
| Hash                 | string                               | ❌       |                                              |
| EnvironmentVariables | map[string]\*string                  | ❌       |                                              |
| Logging              | shared.ContainerLogging              | ❌       |                                              |
| ImageCaching         | bool                                 | ❌       |                                              |

# ContainerLogging

**Properties**

| Name     | Type                    | Required | Description |
| :------- | :---------------------- | :------- | :---------- |
| Axiom    | shared.LoggingAxiom1    | ❌       |             |
| Datadog  | shared.LoggingDatadog1  | ❌       |             |
| NewRelic | shared.LoggingNewRelic1 | ❌       |             |
| Splunk   | shared.LoggingSplunk1   | ❌       |             |
| Tcp      | shared.LoggingTcp1      | ❌       |             |
| Http     | shared.LoggingHttp1     | ❌       |             |

# LoggingAxiom1

**Properties**

| Name     | Type   | Required | Description |
| :------- | :----- | :------- | :---------- |
| Host     | string | ✅       |             |
| ApiToken | string | ✅       |             |
| Dataset  | string | ✅       |             |

# LoggingDatadog1

**Properties**

| Name   | Type                  | Required | Description |
| :----- | :-------------------- | :------- | :---------- |
| Host   | string                | ✅       |             |
| ApiKey | string                | ✅       |             |
| Tags   | []shared.DatadogTags1 | ❌       |             |

# DatadogTags1

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Name  | string | ✅       |             |
| Value | string | ✅       |             |

# LoggingNewRelic1

**Properties**

| Name         | Type   | Required | Description |
| :----------- | :----- | :------- | :---------- |
| Host         | string | ✅       |             |
| IngestionKey | string | ✅       |             |

# LoggingSplunk1

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Host  | string | ✅       |             |
| Token | string | ✅       |             |

# LoggingTcp1

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| Host | string | ✅       |             |
| Port | int64  | ✅       |             |

# LoggingHttp1

**Properties**

| Name        | Type                    | Required | Description |
| :---------- | :---------------------- | :------- | :---------- |
| Host        | string                  | ✅       |             |
| Port        | int64                   | ✅       |             |
| Format      | shared.HttpFormat1      | ✅       |             |
| Compression | shared.HttpCompression1 | ✅       |             |
| User        | string                  | ❌       |             |
| Password    | string                  | ❌       |             |
| Path        | string                  | ❌       |             |
| Headers     | []shared.HttpHeaders1   | ❌       |             |

# HttpFormat1

**Properties**

| Name       | Type   | Required | Description  |
| :--------- | :----- | :------- | :----------- |
| json       | string | ✅       | "json"       |
| json_lines | string | ✅       | "json_lines" |

# HttpCompression1

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| none | string | ✅       | "none"      |
| gzip | string | ✅       | "gzip"      |

# HttpHeaders1

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Name  | string | ✅       |             |
| Value | string | ✅       |             |
