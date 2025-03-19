package workloaderrors

import "encoding/json"

// Represents a workload error
type WorkloadError struct {
	// The timestamp when the workload was initially allocated to a machine
	AllocatedAt *string `json:"allocated_at,omitempty" required:"true"`
	// A detailed error message describing the nature and cause of the workload failure
	Detail *string `json:"detail,omitempty" required:"true" maxLength:"255" minLength:"1" pattern:"^.*$"`
	// The timestamp when the workload failure was detected or reported
	FailedAt *string `json:"failed_at,omitempty" required:"true"`
	// The container group instance identifier.
	InstanceId *string `json:"instance_id,omitempty" required:"true"`
	// The container group machine identifier.
	MachineId *string `json:"machine_id,omitempty" required:"true"`
	// The timestamp when the workload started execution, or null if it failed before starting
	StartedAt *string `json:"started_at,omitempty"`
	// The schema version number for this error record, used for tracking error format changes
	Version *int64 `json:"version,omitempty" required:"true" min:"1" max:"2147483647"`
}

func (w *WorkloadError) GetAllocatedAt() *string {
	if w == nil {
		return nil
	}
	return w.AllocatedAt
}

func (w *WorkloadError) SetAllocatedAt(allocatedAt string) {
	w.AllocatedAt = &allocatedAt
}

func (w *WorkloadError) GetDetail() *string {
	if w == nil {
		return nil
	}
	return w.Detail
}

func (w *WorkloadError) SetDetail(detail string) {
	w.Detail = &detail
}

func (w *WorkloadError) GetFailedAt() *string {
	if w == nil {
		return nil
	}
	return w.FailedAt
}

func (w *WorkloadError) SetFailedAt(failedAt string) {
	w.FailedAt = &failedAt
}

func (w *WorkloadError) GetInstanceId() *string {
	if w == nil {
		return nil
	}
	return w.InstanceId
}

func (w *WorkloadError) SetInstanceId(instanceId string) {
	w.InstanceId = &instanceId
}

func (w *WorkloadError) GetMachineId() *string {
	if w == nil {
		return nil
	}
	return w.MachineId
}

func (w *WorkloadError) SetMachineId(machineId string) {
	w.MachineId = &machineId
}

func (w *WorkloadError) GetStartedAt() *string {
	if w == nil {
		return nil
	}
	return w.StartedAt
}

func (w *WorkloadError) SetStartedAt(startedAt string) {
	w.StartedAt = &startedAt
}

func (w *WorkloadError) GetVersion() *int64 {
	if w == nil {
		return nil
	}
	return w.Version
}

func (w *WorkloadError) SetVersion(version int64) {
	w.Version = &version
}

func (w WorkloadError) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkloadError to string"
	}
	return string(jsonData)
}
