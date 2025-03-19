package shared

import "encoding/json"

// Defines configuration for automatically scaling container instances based on queue length. The autoscaler monitors a queue and adjusts the number of running replicas to maintain the desired queue length.
type QueueBasedAutoscalerConfiguration struct {
	// The target number of items in the queue that the autoscaler attempts to maintain by scaling the containers up or down
	DesiredQueueLength *int64 `json:"desired_queue_length,omitempty" required:"true" min:"1" max:"100"`
	// The maximum number of instances the container can scale up to
	MaxReplicas *int64 `json:"max_replicas,omitempty" required:"true" min:"1" max:"500"`
	// The maximum number of instances that can be removed per minute to prevent rapid downscaling
	MaxDownscalePerMinute *int64 `json:"max_downscale_per_minute,omitempty" min:"1" max:"100"`
	// The maximum number of instances that can be added per minute to prevent rapid upscaling
	MaxUpscalePerMinute *int64 `json:"max_upscale_per_minute,omitempty" min:"1" max:"100"`
	// The minimum number of instances the container can scale down to, ensuring baseline availability
	MinReplicas *int64 `json:"min_replicas,omitempty" required:"true" min:"0" max:"100"`
	// The period (in seconds) in which the autoscaler checks the queue length and applies the scaling formula
	PollingPeriod *int64 `json:"polling_period,omitempty" min:"15" max:"1800"`
}

func (q *QueueBasedAutoscalerConfiguration) GetDesiredQueueLength() *int64 {
	if q == nil {
		return nil
	}
	return q.DesiredQueueLength
}

func (q *QueueBasedAutoscalerConfiguration) SetDesiredQueueLength(desiredQueueLength int64) {
	q.DesiredQueueLength = &desiredQueueLength
}

func (q *QueueBasedAutoscalerConfiguration) GetMaxReplicas() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxReplicas
}

func (q *QueueBasedAutoscalerConfiguration) SetMaxReplicas(maxReplicas int64) {
	q.MaxReplicas = &maxReplicas
}

func (q *QueueBasedAutoscalerConfiguration) GetMaxDownscalePerMinute() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxDownscalePerMinute
}

func (q *QueueBasedAutoscalerConfiguration) SetMaxDownscalePerMinute(maxDownscalePerMinute int64) {
	q.MaxDownscalePerMinute = &maxDownscalePerMinute
}

func (q *QueueBasedAutoscalerConfiguration) GetMaxUpscalePerMinute() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxUpscalePerMinute
}

func (q *QueueBasedAutoscalerConfiguration) SetMaxUpscalePerMinute(maxUpscalePerMinute int64) {
	q.MaxUpscalePerMinute = &maxUpscalePerMinute
}

func (q *QueueBasedAutoscalerConfiguration) GetMinReplicas() *int64 {
	if q == nil {
		return nil
	}
	return q.MinReplicas
}

func (q *QueueBasedAutoscalerConfiguration) SetMinReplicas(minReplicas int64) {
	q.MinReplicas = &minReplicas
}

func (q *QueueBasedAutoscalerConfiguration) GetPollingPeriod() *int64 {
	if q == nil {
		return nil
	}
	return q.PollingPeriod
}

func (q *QueueBasedAutoscalerConfiguration) SetPollingPeriod(pollingPeriod int64) {
	q.PollingPeriod = &pollingPeriod
}

func (q QueueBasedAutoscalerConfiguration) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueBasedAutoscalerConfiguration to string"
	}
	return string(jsonData)
}
