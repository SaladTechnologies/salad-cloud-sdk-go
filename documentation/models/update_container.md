# UpdateContainer

Represents an update container object

**Properties**

| Name                   | Type                                                  | Required | Description                                                                                      |
| :--------------------- | :---------------------------------------------------- | :------- | :----------------------------------------------------------------------------------------------- |
| Image                  | string                                                | ❌       |                                                                                                  |
| Resources              | containergroups.Resources                             | ❌       |                                                                                                  |
| Command                | []string                                              | ❌       | Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image. |
| Priority               | shared.ContainerGroupPriority                         | ❌       |                                                                                                  |
| EnvironmentVariables   | map[string]\*string                                   | ❌       |                                                                                                  |
| Logging                | containergroups.UpdateContainerLogging                | ❌       |                                                                                                  |
| RegistryAuthentication | containergroups.UpdateContainerRegistryAuthentication | ❌       |                                                                                                  |

# Resources

**Properties**

| Name          | Type     | Required | Description |
| :------------ | :------- | :------- | :---------- |
| Cpu           | int64    | ❌       |             |
| Memory        | int64    | ❌       |             |
| GpuClasses    | []string | ❌       |             |
| StorageAmount | int64    | ❌       |             |

# UpdateContainerLogging

**Properties**

| Name     | Type                             | Required | Description |
| :------- | :------------------------------- | :------- | :---------- |
| Axiom    | containergroups.LoggingAxiom3    | ❌       |             |
| Datadog  | containergroups.LoggingDatadog3  | ❌       |             |
| NewRelic | containergroups.LoggingNewRelic3 | ❌       |             |
| Splunk   | containergroups.LoggingSplunk3   | ❌       |             |
| Tcp      | containergroups.LoggingTcp3      | ❌       |             |
| Http     | containergroups.LoggingHttp3     | ❌       |             |

# LoggingAxiom3

**Properties**

| Name     | Type   | Required | Description |
| :------- | :----- | :------- | :---------- |
| Host     | string | ✅       |             |
| ApiToken | string | ✅       |             |
| Dataset  | string | ✅       |             |

# LoggingDatadog3

**Properties**

| Name   | Type                           | Required | Description |
| :----- | :----------------------------- | :------- | :---------- |
| Host   | string                         | ✅       |             |
| ApiKey | string                         | ✅       |             |
| Tags   | []containergroups.DatadogTags3 | ❌       |             |

# DatadogTags3

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Name  | string | ✅       |             |
| Value | string | ✅       |             |

# LoggingNewRelic3

**Properties**

| Name         | Type   | Required | Description |
| :----------- | :----- | :------- | :---------- |
| Host         | string | ✅       |             |
| IngestionKey | string | ✅       |             |

# LoggingSplunk3

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Host  | string | ✅       |             |
| Token | string | ✅       |             |

# LoggingTcp3

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| Host | string | ✅       |             |
| Port | int64  | ✅       |             |

# LoggingHttp3

**Properties**

| Name        | Type                             | Required | Description |
| :---------- | :------------------------------- | :------- | :---------- |
| Host        | string                           | ✅       |             |
| Port        | int64                            | ✅       |             |
| Format      | containergroups.HttpFormat3      | ✅       |             |
| Compression | containergroups.HttpCompression3 | ✅       |             |
| User        | string                           | ❌       |             |
| Password    | string                           | ❌       |             |
| Path        | string                           | ❌       |             |
| Headers     | []containergroups.HttpHeaders4   | ❌       |             |

# HttpFormat3

**Properties**

| Name       | Type   | Required | Description  |
| :--------- | :----- | :------- | :----------- |
| json       | string | ✅       | "json"       |
| json_lines | string | ✅       | "json_lines" |

# HttpCompression3

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| none | string | ✅       | "none"      |
| gzip | string | ✅       | "gzip"      |

# HttpHeaders4

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Name  | string | ✅       |             |
| Value | string | ✅       |             |

# UpdateContainerRegistryAuthentication

**Properties**

| Name      | Type                                             | Required | Description |
| :-------- | :----------------------------------------------- | :------- | :---------- |
| Basic     | containergroups.RegistryAuthenticationBasic2     | ❌       |             |
| GcpGcr    | containergroups.RegistryAuthenticationGcpGcr2    | ❌       |             |
| AwsEcr    | containergroups.RegistryAuthenticationAwsEcr2    | ❌       |             |
| DockerHub | containergroups.RegistryAuthenticationDockerHub2 | ❌       |             |
| GcpGar    | containergroups.RegistryAuthenticationGcpGar2    | ❌       |             |

# RegistryAuthenticationBasic2

**Properties**

| Name     | Type   | Required | Description |
| :------- | :----- | :------- | :---------- |
| Username | string | ✅       |             |
| Password | string | ✅       |             |

# RegistryAuthenticationGcpGcr2

**Properties**

| Name       | Type   | Required | Description |
| :--------- | :----- | :------- | :---------- |
| ServiceKey | string | ✅       |             |

# RegistryAuthenticationAwsEcr2

**Properties**

| Name            | Type   | Required | Description |
| :-------------- | :----- | :------- | :---------- |
| AccessKeyId     | string | ✅       |             |
| SecretAccessKey | string | ✅       |             |

# RegistryAuthenticationDockerHub2

**Properties**

| Name                | Type   | Required | Description |
| :------------------ | :----- | :------- | :---------- |
| Username            | string | ✅       |             |
| PersonalAccessToken | string | ✅       |             |

# RegistryAuthenticationGcpGar2

**Properties**

| Name       | Type   | Required | Description |
| :--------- | :----- | :------- | :---------- |
| ServiceKey | string | ✅       |             |
