# ContainerRegistryAuthentication

Authentication configuration for various container registry types, including AWS ECR, Docker Hub, GCP GAR, GCP GCR, and basic authentication.

**Properties**

| Name      | Type                                                     | Required | Description                                                                 |
| :-------- | :------------------------------------------------------- | :------- | :-------------------------------------------------------------------------- |
| AwsEcr    | containergroups.ContainerRegistryAuthenticationAwsEcr    | ❌       | Authentication details for AWS Elastic Container Registry (ECR)             |
| Basic     | containergroups.ContainerRegistryAuthenticationBasic     | ❌       | Basic username and password authentication for generic container registries |
| DockerHub | containergroups.ContainerRegistryAuthenticationDockerHub | ❌       | Authentication details for Docker Hub registry                              |
| GcpGar    | containergroups.ContainerRegistryAuthenticationGcpGar    | ❌       | Authentication details for Google Artifact Registry (GAR)                   |
| GcpGcr    | containergroups.ContainerRegistryAuthenticationGcpGcr    | ❌       | Authentication details for Google Container Registry (GCR)                  |
