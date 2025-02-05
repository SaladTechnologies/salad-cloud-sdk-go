package shared

import (
	"encoding/json"
)

// Represents the autoscaling rules for a queue
type QueueAutoscaler struct {
	// The minimum number of instances the container can scale down to
	MinReplicas *int64 `json:"min_replicas,omitempty" required:"true" min:"0" max:"100"`
	// The maximum number of instances the container can scale up to
	MaxReplicas        *int64 `json:"max_replicas,omitempty" required:"true" min:"1" max:"250"`
	DesiredQueueLength *int64 `json:"desired_queue_length,omitempty" required:"true" min:"1" max:"100"`
	// The period (in seconds) in which the queue checks the formula
	PollingPeriod *int64 `json:"polling_period,omitempty" min:"15" max:"1800"`
	// The maximum number of instances that can be added per minute
	MaxUpscalePerMinute *int64 `json:"max_upscale_per_minute,omitempty" min:"1" max:"100"`
	// The maximum number of instances that can be removed per minute
	MaxDownscalePerMinute *int64 `json:"max_downscale_per_minute,omitempty" min:"1" max:"100"`
	touched               map[string]bool
}

func (q *QueueAutoscaler) GetMinReplicas() *int64 {
	if q == nil {
		return nil
	}
	return q.MinReplicas
}

func (q *QueueAutoscaler) SetMinReplicas(minReplicas int64) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["MinReplicas"] = true
	q.MinReplicas = &minReplicas
}

func (q *QueueAutoscaler) SetMinReplicasNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["MinReplicas"] = true
	q.MinReplicas = nil
}

func (q *QueueAutoscaler) GetMaxReplicas() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxReplicas
}

func (q *QueueAutoscaler) SetMaxReplicas(maxReplicas int64) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["MaxReplicas"] = true
	q.MaxReplicas = &maxReplicas
}

func (q *QueueAutoscaler) SetMaxReplicasNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["MaxReplicas"] = true
	q.MaxReplicas = nil
}

func (q *QueueAutoscaler) GetDesiredQueueLength() *int64 {
	if q == nil {
		return nil
	}
	return q.DesiredQueueLength
}

func (q *QueueAutoscaler) SetDesiredQueueLength(desiredQueueLength int64) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["DesiredQueueLength"] = true
	q.DesiredQueueLength = &desiredQueueLength
}

func (q *QueueAutoscaler) SetDesiredQueueLengthNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["DesiredQueueLength"] = true
	q.DesiredQueueLength = nil
}

func (q *QueueAutoscaler) GetPollingPeriod() *int64 {
	if q == nil {
		return nil
	}
	return q.PollingPeriod
}

func (q *QueueAutoscaler) SetPollingPeriod(pollingPeriod int64) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["PollingPeriod"] = true
	q.PollingPeriod = &pollingPeriod
}

func (q *QueueAutoscaler) SetPollingPeriodNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["PollingPeriod"] = true
	q.PollingPeriod = nil
}

func (q *QueueAutoscaler) GetMaxUpscalePerMinute() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxUpscalePerMinute
}

func (q *QueueAutoscaler) SetMaxUpscalePerMinute(maxUpscalePerMinute int64) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["MaxUpscalePerMinute"] = true
	q.MaxUpscalePerMinute = &maxUpscalePerMinute
}

func (q *QueueAutoscaler) SetMaxUpscalePerMinuteNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["MaxUpscalePerMinute"] = true
	q.MaxUpscalePerMinute = nil
}

func (q *QueueAutoscaler) GetMaxDownscalePerMinute() *int64 {
	if q == nil {
		return nil
	}
	return q.MaxDownscalePerMinute
}

func (q *QueueAutoscaler) SetMaxDownscalePerMinute(maxDownscalePerMinute int64) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["MaxDownscalePerMinute"] = true
	q.MaxDownscalePerMinute = &maxDownscalePerMinute
}

func (q *QueueAutoscaler) SetMaxDownscalePerMinuteNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["MaxDownscalePerMinute"] = true
	q.MaxDownscalePerMinute = nil
}

func (q QueueAutoscaler) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if q.touched["MinReplicas"] && q.MinReplicas == nil {
		data["min_replicas"] = nil
	} else if q.MinReplicas != nil {
		data["min_replicas"] = q.MinReplicas
	}

	if q.touched["MaxReplicas"] && q.MaxReplicas == nil {
		data["max_replicas"] = nil
	} else if q.MaxReplicas != nil {
		data["max_replicas"] = q.MaxReplicas
	}

	if q.touched["DesiredQueueLength"] && q.DesiredQueueLength == nil {
		data["desired_queue_length"] = nil
	} else if q.DesiredQueueLength != nil {
		data["desired_queue_length"] = q.DesiredQueueLength
	}

	if q.touched["PollingPeriod"] && q.PollingPeriod == nil {
		data["polling_period"] = nil
	} else if q.PollingPeriod != nil {
		data["polling_period"] = q.PollingPeriod
	}

	if q.touched["MaxUpscalePerMinute"] && q.MaxUpscalePerMinute == nil {
		data["max_upscale_per_minute"] = nil
	} else if q.MaxUpscalePerMinute != nil {
		data["max_upscale_per_minute"] = q.MaxUpscalePerMinute
	}

	if q.touched["MaxDownscalePerMinute"] && q.MaxDownscalePerMinute == nil {
		data["max_downscale_per_minute"] = nil
	} else if q.MaxDownscalePerMinute != nil {
		data["max_downscale_per_minute"] = q.MaxDownscalePerMinute
	}

	return json.Marshal(data)
}

func (q QueueAutoscaler) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueAutoscaler to string"
	}
	return string(jsonData)
}
