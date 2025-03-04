# InferenceEndpointList

Represents a page from the collection of inference endpoints.

**Properties**

| Name      | Type                                   | Required | Description                                  |
| :-------- | :------------------------------------- | :------- | :------------------------------------------- |
| Items     | []inferenceendpoints.InferenceEndpoint | ✅       | The list of inference endpoints.             |
| Page      | int64                                  | ✅       | The page number.                             |
| PageSize  | int64                                  | ✅       | The maximum number of items per page.        |
| TotalSize | int64                                  | ✅       | The total number of items in the collection. |
