# ContainerGroupNetworking

Represents container group networking parameters

**Properties**

| Name                  | Type                                        | Required | Description |
| :-------------------- | :------------------------------------------ | :------- | :---------- |
| Protocol              | shared.ContainerNetworkingProtocol          | ✅       |             |
| Port                  | int64                                       | ✅       |             |
| Auth                  | bool                                        | ✅       |             |
| Dns                   | string                                      | ✅       |             |
| LoadBalancer          | shared.ContainerGroupNetworkingLoadBalancer | ❌       |             |
| SingleConnectionLimit | bool                                        | ❌       |             |
| ClientRequestTimeout  | int64                                       | ❌       |             |
| ServerResponseTimeout | int64                                       | ❌       |             |

# ContainerGroupNetworkingLoadBalancer

**Properties**

| Name                        | Type   | Required | Description                   |
| :-------------------------- | :----- | :------- | :---------------------------- |
| round_robin                 | string | ✅       | "round_robin"                 |
| least_number_of_connections | string | ✅       | "least_number_of_connections" |
