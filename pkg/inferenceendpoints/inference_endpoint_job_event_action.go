package inferenceendpoints

type InferenceEndpointJobEventAction string

const (
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_CREATED   InferenceEndpointJobEventAction = "created"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_STARTED   InferenceEndpointJobEventAction = "started"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_SUCCEEDED InferenceEndpointJobEventAction = "succeeded"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_CANCELLED InferenceEndpointJobEventAction = "cancelled"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_FAILED    InferenceEndpointJobEventAction = "failed"
)
