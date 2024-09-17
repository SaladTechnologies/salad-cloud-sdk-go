package shared

// Represents the autoscaling rules for a queue
type QueueAutoscaler struct {
	MinReplicas           *int64 `json:"min_replicas,omitempty" required:"true" min:"0" max:"100"`
	MaxReplicas           *int64 `json:"max_replicas,omitempty" required:"true" min:"1" max:"250"`
	DesiredQueueLength    *int64 `json:"desired_queue_length,omitempty" required:"true" min:"1" max:"100"`
	PollingPeriod         *int64 `json:"polling_period,omitempty" min:"15" max:"1800"`
	MaxUpscalePerMinute   *int64 `json:"max_upscale_per_minute,omitempty" min:"1" max:"100"`
	MaxDownscalePerMinute *int64 `json:"max_downscale_per_minute,omitempty" min:"1" max:"100"`
}

func (q *QueueAutoscaler) SetMinReplicas(minReplicas int64) {
	q.MinReplicas = &minReplicas
}

func (q *QueueAutoscaler) GetMinReplicas() *int64 {
	if q == nil {
		return nil
	}
	return q.MinReplicas
}

func (q *QueueAutoscaler) SetMaxReplicas(maxReplicas int64) {
	q.MaxReplicas = &maxReplicas
}

func (q *QueueAutoscaler) GetMaxReplicas() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxReplicas
}

func (q *QueueAutoscaler) SetDesiredQueueLength(desiredQueueLength int64) {
	q.DesiredQueueLength = &desiredQueueLength
}

func (q *QueueAutoscaler) GetDesiredQueueLength() *int64 {
	if q == nil {
		return nil
	}
	return q.DesiredQueueLength
}

func (q *QueueAutoscaler) SetPollingPeriod(pollingPeriod int64) {
	q.PollingPeriod = &pollingPeriod
}

func (q *QueueAutoscaler) GetPollingPeriod() *int64 {
	if q == nil {
		return nil
	}
	return q.PollingPeriod
}

func (q *QueueAutoscaler) SetMaxUpscalePerMinute(maxUpscalePerMinute int64) {
	q.MaxUpscalePerMinute = &maxUpscalePerMinute
}

func (q *QueueAutoscaler) GetMaxUpscalePerMinute() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxUpscalePerMinute
}

func (q *QueueAutoscaler) SetMaxDownscalePerMinute(maxDownscalePerMinute int64) {
	q.MaxDownscalePerMinute = &maxDownscalePerMinute
}

func (q *QueueAutoscaler) GetMaxDownscalePerMinute() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxDownscalePerMinute
}
