package shared

import "encoding/json"

// Represents the autoscaling rules for a queue
type QueueAutoscaler struct {
	// The minimum number of instances the container can scale down to
	MinReplicas *int64 `json:"min_replicas,omitempty" required:"true" min:"0" max:"100"`
	// The maximum number of instances the container can scale up to
	MaxReplicas        *int64 `json:"max_replicas,omitempty" required:"true" min:"1" max:"500"`
	DesiredQueueLength *int64 `json:"desired_queue_length,omitempty" required:"true" min:"1" max:"100"`
	// The period (in seconds) in which the queue checks the formula
	PollingPeriod *int64 `json:"polling_period,omitempty" min:"15" max:"1800"`
	// The maximum number of instances that can be added per minute
	MaxUpscalePerMinute *int64 `json:"max_upscale_per_minute,omitempty" min:"1" max:"100"`
	// The maximum number of instances that can be removed per minute
	MaxDownscalePerMinute *int64 `json:"max_downscale_per_minute,omitempty" min:"1" max:"100"`
}

func (q *QueueAutoscaler) GetMinReplicas() *int64 {
	if q == nil {
		return nil
	}
	return q.MinReplicas
}

func (q *QueueAutoscaler) SetMinReplicas(minReplicas int64) {
	q.MinReplicas = &minReplicas
}

func (q *QueueAutoscaler) GetMaxReplicas() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxReplicas
}

func (q *QueueAutoscaler) SetMaxReplicas(maxReplicas int64) {
	q.MaxReplicas = &maxReplicas
}

func (q *QueueAutoscaler) GetDesiredQueueLength() *int64 {
	if q == nil {
		return nil
	}
	return q.DesiredQueueLength
}

func (q *QueueAutoscaler) SetDesiredQueueLength(desiredQueueLength int64) {
	q.DesiredQueueLength = &desiredQueueLength
}

func (q *QueueAutoscaler) GetPollingPeriod() *int64 {
	if q == nil {
		return nil
	}
	return q.PollingPeriod
}

func (q *QueueAutoscaler) SetPollingPeriod(pollingPeriod int64) {
	q.PollingPeriod = &pollingPeriod
}

func (q *QueueAutoscaler) GetMaxUpscalePerMinute() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxUpscalePerMinute
}

func (q *QueueAutoscaler) SetMaxUpscalePerMinute(maxUpscalePerMinute int64) {
	q.MaxUpscalePerMinute = &maxUpscalePerMinute
}

func (q *QueueAutoscaler) GetMaxDownscalePerMinute() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxDownscalePerMinute
}

func (q *QueueAutoscaler) SetMaxDownscalePerMinute(maxDownscalePerMinute int64) {
	q.MaxDownscalePerMinute = &maxDownscalePerMinute
}

func (q QueueAutoscaler) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueAutoscaler to string"
	}
	return string(jsonData)
}
