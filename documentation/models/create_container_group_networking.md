# CreateContainerGroupNetworking

Network configuration for container groups specifying connectivity parameters, including authentication, protocol, and timeout settings

**Properties**

| Name                  | Type                                                                                              | Required | Description                                                                                                                                                       |
| :-------------------- | :------------------------------------------------------------------------------------------------ | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Auth                  | bool                                                                                              | ✅       | Determines whether authentication is required for network connections to the container group                                                                      |
| Port                  | int64                                                                                             | ✅       | The container group networking port.                                                                                                                              |
| Protocol              | [shared.ContainerNetworkingProtocol](container_networking_protocol.md)                            | ✅       | Defines the communication protocol used for network traffic between containers or external systems. Currently supports HTTP protocol for web-based communication. |
| ClientRequestTimeout  | int64                                                                                             | ❌       | The container group networking client request timeout.                                                                                                            |
| LoadBalancer          | [shared.TheContainerGroupNetworkingLoadBalancer](the_container_group_networking_load_balancer.md) | ❌       | The container group networking load balancer.                                                                                                                     |
| ServerResponseTimeout | int64                                                                                             | ❌       | The container group networking server response timeout.                                                                                                           |
| SingleConnectionLimit | bool                                                                                              | ❌       | The container group networking single connection limit flag.                                                                                                      |
