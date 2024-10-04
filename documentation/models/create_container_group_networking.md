# CreateContainerGroupNetworking

Represents container group networking parameters

**Properties**

| Name                  | Type                                                       | Required | Description |
| :-------------------- | :--------------------------------------------------------- | :------- | :---------- |
| Protocol              | shared.ContainerNetworkingProtocol                         | ✅       |             |
| Port                  | int64                                                      | ✅       |             |
| Auth                  | bool                                                       | ✅       |             |
| LoadBalancer          | containergroups.CreateContainerGroupNetworkingLoadBalancer | ❌       |             |
| SingleConnectionLimit | bool                                                       | ❌       |             |
| ClientRequestTimeout  | int64                                                      | ❌       |             |
| ServerResponseTimeout | int64                                                      | ❌       |             |

# CreateContainerGroupNetworkingLoadBalancer

**Properties**

| Name                        | Type   | Required | Description                   |
| :-------------------------- | :----- | :------- | :---------------------------- |
| round_robin                 | string | ✅       | "round_robin"                 |
| least_number_of_connections | string | ✅       | "least_number_of_connections" |
