package shared

// The container group networking load balancer.
type TheContainerGroupNetworkingLoadBalancer string

const (
	THE_CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_ROUND_ROBIN                 TheContainerGroupNetworkingLoadBalancer = "round_robin"
	THE_CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_LEAST_NUMBER_OF_CONNECTIONS TheContainerGroupNetworkingLoadBalancer = "least_number_of_connections"
)
