# ContainerGroupStatus

Represents the current operational state of a container group within the Salad platform.

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| pending   | string | ✅       | "pending"   |
| running   | string | ✅       | "running"   |
| stopped   | string | ✅       | "stopped"   |
| succeeded | string | ✅       | "succeeded" |
| failed    | string | ✅       | "failed"    |
| deploying | string | ✅       | "deploying" |
