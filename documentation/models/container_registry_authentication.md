# ContainerRegistryAuthentication

Authentication configuration for various container registry types, including AWS ECR, Docker Hub, GCP GAR, GCP GCR, and basic authentication.

**Properties**

| Name      | Type                                                                                                        | Required | Description                                                                 |
| :-------- | :---------------------------------------------------------------------------------------------------------- | :------- | :-------------------------------------------------------------------------- |
| AwsEcr    | [containergroups.ContainerRegistryAuthenticationAwsEcr](container_registry_authentication_aws_ecr.md)       | ❌       | Authentication details for AWS Elastic Container Registry (ECR)             |
| Basic     | [containergroups.ContainerRegistryAuthenticationBasic](container_registry_authentication_basic.md)          | ❌       | Basic username and password authentication for generic container registries |
| DockerHub | [containergroups.ContainerRegistryAuthenticationDockerHub](container_registry_authentication_docker_hub.md) | ❌       | Authentication details for Docker Hub registry                              |
| GcpGar    | [containergroups.ContainerRegistryAuthenticationGcpGar](container_registry_authentication_gcp_gar.md)       | ❌       | Authentication details for Google Artifact Registry (GAR)                   |
| GcpGcr    | [containergroups.ContainerRegistryAuthenticationGcpGcr](container_registry_authentication_gcp_gcr.md)       | ❌       | Authentication details for Google Container Registry (GCR)                  |
