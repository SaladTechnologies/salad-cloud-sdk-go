# QueueJobPrototype

Represents a request to create a queue job

**Properties**

| Name     | Type   | Required | Description                                |
| :------- | :----- | :------- | :----------------------------------------- |
| Input    | any    | ✅       | The job input. May be any valid JSON.      |
| Metadata | any    | ❌       | Additional metadata for the job            |
| Webhook  | string | ❌       | The webhook to call when the job completes |
