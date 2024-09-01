# CreateContainer

Represents a container

**Properties**

| Name                   | Type                                                  | Required | Description                                                                                      |
| :--------------------- | :---------------------------------------------------- | :------- | :----------------------------------------------------------------------------------------------- |
| Image                  | string                                                | ✅       |                                                                                                  |
| Resources              | shared.ContainerResourceRequirements                  | ✅       | Represents a container resource requirements                                                     |
| Command                | []string                                              | ❌       | Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image. |
| Priority               | shared.ContainerGroupPriority                         | ❌       |                                                                                                  |
| EnvironmentVariables   | map[string]\*string                                   | ❌       |                                                                                                  |
| Logging                | containergroups.CreateContainerLogging                | ❌       |                                                                                                  |
| RegistryAuthentication | containergroups.CreateContainerRegistryAuthentication | ❌       |                                                                                                  |

# CreateContainerLogging

**Properties**

| Name     | Type                             | Required | Description |
| :------- | :------------------------------- | :------- | :---------- |
| Axiom    | containergroups.LoggingAxiom2    | ❌       |             |
| Datadog  | containergroups.LoggingDatadog2  | ❌       |             |
| NewRelic | containergroups.LoggingNewRelic2 | ❌       |             |
| Splunk   | containergroups.LoggingSplunk2   | ❌       |             |
| Tcp      | containergroups.LoggingTcp2      | ❌       |             |
| Http     | containergroups.LoggingHttp2     | ❌       |             |

# LoggingAxiom2

**Properties**

| Name     | Type   | Required | Description |
| :------- | :----- | :------- | :---------- |
| Host     | string | ✅       |             |
| ApiToken | string | ✅       |             |
| Dataset  | string | ✅       |             |

# LoggingDatadog2

**Properties**

| Name   | Type                           | Required | Description |
| :----- | :----------------------------- | :------- | :---------- |
| Host   | string                         | ✅       |             |
| ApiKey | string                         | ✅       |             |
| Tags   | []containergroups.DatadogTags2 | ❌       |             |

# DatadogTags2

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Name  | string | ✅       |             |
| Value | string | ✅       |             |

# LoggingNewRelic2

**Properties**

| Name         | Type   | Required | Description |
| :----------- | :----- | :------- | :---------- |
| Host         | string | ✅       |             |
| IngestionKey | string | ✅       |             |

# LoggingSplunk2

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Host  | string | ✅       |             |
| Token | string | ✅       |             |

# LoggingTcp2

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| Host | string | ✅       |             |
| Port | int64  | ✅       |             |

# LoggingHttp2

**Properties**

| Name        | Type                             | Required | Description |
| :---------- | :------------------------------- | :------- | :---------- |
| Host        | string                           | ✅       |             |
| Port        | int64                            | ✅       |             |
| Format      | containergroups.HttpFormat2      | ✅       |             |
| Compression | containergroups.HttpCompression2 | ✅       |             |
| User        | string                           | ❌       |             |
| Password    | string                           | ❌       |             |
| Path        | string                           | ❌       |             |
| Headers     | []containergroups.HttpHeaders3   | ❌       |             |

# HttpFormat2

**Properties**

| Name       | Type   | Required | Description  |
| :--------- | :----- | :------- | :----------- |
| json       | string | ✅       | "json"       |
| json_lines | string | ✅       | "json_lines" |

# HttpCompression2

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| none | string | ✅       | "none"      |
| gzip | string | ✅       | "gzip"      |

# HttpHeaders3

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Name  | string | ✅       |             |
| Value | string | ✅       |             |

# CreateContainerRegistryAuthentication

**Properties**

| Name      | Type                                             | Required | Description |
| :-------- | :----------------------------------------------- | :------- | :---------- |
| Basic     | containergroups.RegistryAuthenticationBasic1     | ❌       |             |
| GcpGcr    | containergroups.RegistryAuthenticationGcpGcr1    | ❌       |             |
| AwsEcr    | containergroups.RegistryAuthenticationAwsEcr1    | ❌       |             |
| DockerHub | containergroups.RegistryAuthenticationDockerHub1 | ❌       |             |
| GcpGar    | containergroups.RegistryAuthenticationGcpGar1    | ❌       |             |

# RegistryAuthenticationBasic1

**Properties**

| Name     | Type   | Required | Description |
| :------- | :----- | :------- | :---------- |
| Username | string | ✅       |             |
| Password | string | ✅       |             |

# RegistryAuthenticationGcpGcr1

**Properties**

| Name       | Type   | Required | Description |
| :--------- | :----- | :------- | :---------- |
| ServiceKey | string | ✅       |             |

# RegistryAuthenticationAwsEcr1

**Properties**

| Name            | Type   | Required | Description |
| :-------------- | :----- | :------- | :---------- |
| AccessKeyId     | string | ✅       |             |
| SecretAccessKey | string | ✅       |             |

# RegistryAuthenticationDockerHub1

**Properties**

| Name                | Type   | Required | Description |
| :------------------ | :----- | :------- | :---------- |
| Username            | string | ✅       |             |
| PersonalAccessToken | string | ✅       |             |

# RegistryAuthenticationGcpGar1

**Properties**

| Name       | Type   | Required | Description |
| :--------- | :----- | :------- | :---------- |
| ServiceKey | string | ✅       |             |
