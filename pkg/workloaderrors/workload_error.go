package workloaderrors

// Represents a workload error
type WorkloadError struct {
	Detail      *string `json:"detail,omitempty" required:"true"`
	FailedAt    *string `json:"failed_at,omitempty" required:"true"`
	InstanceId  *string `json:"instance_id,omitempty" required:"true"`
	MachineId   *string `json:"machine_id,omitempty" required:"true"`
	AllocatedAt *string `json:"allocated_at,omitempty" required:"true"`
	StartedAt   *string `json:"started_at,omitempty"`
	Version     *int64  `json:"version,omitempty" required:"true" min:"1"`
}

func (w *WorkloadError) SetDetail(detail string) {
	w.Detail = &detail
}

func (w *WorkloadError) GetDetail() *string {
	if w == nil {
		return nil
	}
	return w.Detail
}

func (w *WorkloadError) SetFailedAt(failedAt string) {
	w.FailedAt = &failedAt
}

func (w *WorkloadError) GetFailedAt() *string {
	if w == nil {
		return nil
	}
	return w.FailedAt
}

func (w *WorkloadError) SetInstanceId(instanceId string) {
	w.InstanceId = &instanceId
}

func (w *WorkloadError) GetInstanceId() *string {
	if w == nil {
		return nil
	}
	return w.InstanceId
}

func (w *WorkloadError) SetMachineId(machineId string) {
	w.MachineId = &machineId
}

func (w *WorkloadError) GetMachineId() *string {
	if w == nil {
		return nil
	}
	return w.MachineId
}

func (w *WorkloadError) SetAllocatedAt(allocatedAt string) {
	w.AllocatedAt = &allocatedAt
}

func (w *WorkloadError) GetAllocatedAt() *string {
	if w == nil {
		return nil
	}
	return w.AllocatedAt
}

func (w *WorkloadError) SetStartedAt(startedAt string) {
	w.StartedAt = &startedAt
}

func (w *WorkloadError) GetStartedAt() *string {
	if w == nil {
		return nil
	}
	return w.StartedAt
}

func (w *WorkloadError) SetVersion(version int64) {
	w.Version = &version
}

func (w *WorkloadError) GetVersion() *int64 {
	if w == nil {
		return nil
	}
	return w.Version
}
