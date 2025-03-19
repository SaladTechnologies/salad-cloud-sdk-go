package shared

// Defines the communication protocol used for network traffic between containers or external systems. Currently supports HTTP protocol for web-based communication.
type ContainerNetworkingProtocol string

const (
	CONTAINER_NETWORKING_PROTOCOL_HTTP ContainerNetworkingProtocol = "http"
)
