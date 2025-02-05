package workloaderrors

import (
	"encoding/json"
)

// Represents a workload error
type WorkloadError struct {
	Detail      *string `json:"detail,omitempty" required:"true"`
	FailedAt    *string `json:"failed_at,omitempty" required:"true"`
	InstanceId  *string `json:"instance_id,omitempty" required:"true"`
	MachineId   *string `json:"machine_id,omitempty" required:"true"`
	AllocatedAt *string `json:"allocated_at,omitempty" required:"true"`
	StartedAt   *string `json:"started_at,omitempty"`
	Version     *int64  `json:"version,omitempty" required:"true" min:"1"`
	touched     map[string]bool
}

func (w *WorkloadError) GetDetail() *string {
	if w == nil {
		return nil
	}
	return w.Detail
}

func (w *WorkloadError) SetDetail(detail string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Detail"] = true
	w.Detail = &detail
}

func (w *WorkloadError) SetDetailNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Detail"] = true
	w.Detail = nil
}

func (w *WorkloadError) GetFailedAt() *string {
	if w == nil {
		return nil
	}
	return w.FailedAt
}

func (w *WorkloadError) SetFailedAt(failedAt string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["FailedAt"] = true
	w.FailedAt = &failedAt
}

func (w *WorkloadError) SetFailedAtNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["FailedAt"] = true
	w.FailedAt = nil
}

func (w *WorkloadError) GetInstanceId() *string {
	if w == nil {
		return nil
	}
	return w.InstanceId
}

func (w *WorkloadError) SetInstanceId(instanceId string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["InstanceId"] = true
	w.InstanceId = &instanceId
}

func (w *WorkloadError) SetInstanceIdNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["InstanceId"] = true
	w.InstanceId = nil
}

func (w *WorkloadError) GetMachineId() *string {
	if w == nil {
		return nil
	}
	return w.MachineId
}

func (w *WorkloadError) SetMachineId(machineId string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["MachineId"] = true
	w.MachineId = &machineId
}

func (w *WorkloadError) SetMachineIdNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["MachineId"] = true
	w.MachineId = nil
}

func (w *WorkloadError) GetAllocatedAt() *string {
	if w == nil {
		return nil
	}
	return w.AllocatedAt
}

func (w *WorkloadError) SetAllocatedAt(allocatedAt string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["AllocatedAt"] = true
	w.AllocatedAt = &allocatedAt
}

func (w *WorkloadError) SetAllocatedAtNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["AllocatedAt"] = true
	w.AllocatedAt = nil
}

func (w *WorkloadError) GetStartedAt() *string {
	if w == nil {
		return nil
	}
	return w.StartedAt
}

func (w *WorkloadError) SetStartedAt(startedAt string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["StartedAt"] = true
	w.StartedAt = &startedAt
}

func (w *WorkloadError) SetStartedAtNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["StartedAt"] = true
	w.StartedAt = nil
}

func (w *WorkloadError) GetVersion() *int64 {
	if w == nil {
		return nil
	}
	return w.Version
}

func (w *WorkloadError) SetVersion(version int64) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Version"] = true
	w.Version = &version
}

func (w *WorkloadError) SetVersionNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Version"] = true
	w.Version = nil
}

func (w WorkloadError) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if w.touched["Detail"] && w.Detail == nil {
		data["detail"] = nil
	} else if w.Detail != nil {
		data["detail"] = w.Detail
	}

	if w.touched["FailedAt"] && w.FailedAt == nil {
		data["failed_at"] = nil
	} else if w.FailedAt != nil {
		data["failed_at"] = w.FailedAt
	}

	if w.touched["InstanceId"] && w.InstanceId == nil {
		data["instance_id"] = nil
	} else if w.InstanceId != nil {
		data["instance_id"] = w.InstanceId
	}

	if w.touched["MachineId"] && w.MachineId == nil {
		data["machine_id"] = nil
	} else if w.MachineId != nil {
		data["machine_id"] = w.MachineId
	}

	if w.touched["AllocatedAt"] && w.AllocatedAt == nil {
		data["allocated_at"] = nil
	} else if w.AllocatedAt != nil {
		data["allocated_at"] = w.AllocatedAt
	}

	if w.touched["StartedAt"] && w.StartedAt == nil {
		data["started_at"] = nil
	} else if w.StartedAt != nil {
		data["started_at"] = w.StartedAt
	}

	if w.touched["Version"] && w.Version == nil {
		data["version"] = nil
	} else if w.Version != nil {
		data["version"] = w.Version
	}

	return json.Marshal(data)
}

func (w WorkloadError) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkloadError to string"
	}
	return string(jsonData)
}
