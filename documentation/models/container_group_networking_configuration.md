# ContainerGroupNetworkingConfiguration

Network configuration for container groups that defines connectivity, routing, and access control settings

**Properties**

| Name                  | Type                                                                                              | Required | Description                                                                                                                                                       |
| :-------------------- | :------------------------------------------------------------------------------------------------ | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Auth                  | bool                                                                                              | ✅       | Whether authentication is required for network access to the container group                                                                                      |
| Dns                   | string                                                                                            | ✅       | Domain name or URL endpoint for the container group's network interface                                                                                           |
| LoadBalancer          | [shared.TheContainerGroupNetworkingLoadBalancer](the_container_group_networking_load_balancer.md) | ✅       | The container group networking load balancer.                                                                                                                     |
| Port                  | int64                                                                                             | ✅       | The container group networking port.                                                                                                                              |
| Protocol              | [shared.ContainerNetworkingProtocol](container_networking_protocol.md)                            | ✅       | Defines the communication protocol used for network traffic between containers or external systems. Currently supports HTTP protocol for web-based communication. |
| ClientRequestTimeout  | int64                                                                                             | ❌       | The container group networking client request timeout.                                                                                                            |
| ServerResponseTimeout | int64                                                                                             | ❌       | The container group networking server response timeout.                                                                                                           |
| SingleConnectionLimit | bool                                                                                              | ❌       | The container group networking single connection limit flag.                                                                                                      |
