package systemlogs

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a system log
type SystemLog struct {
	// The name of the event
	EventName *string `json:"event_name,omitempty" required:"true"`
	// The UTC date & time when the log item was created
	EventTime *string `json:"event_time,omitempty" required:"true"`
	// The unique instance ID
	InstanceId *string `json:"instance_id,omitempty"`
	// The organization-specific machine ID
	MachineId *string `json:"machine_id,omitempty"`
	// The version instance ID
	Version *string `json:"version,omitempty" required:"true"`
	// The number of CPUs
	ResourceCpu *util.Nullable[int64] `json:"resource_cpu,omitempty" required:"true" min:"1" max:"16"`
	// The memory amount in MB
	ResourceMemory *util.Nullable[int64] `json:"resource_memory,omitempty" required:"true" min:"1024" max:"61440"`
	// The GPU class name
	ResourceGpuClass *string `json:"resource_gpu_class,omitempty" required:"true"`
	// The storage amount in bytes
	ResourceStorageAmount *util.Nullable[int64] `json:"resource_storage_amount,omitempty" required:"true" min:"1073741824" max:"53687091200"`
}

func (s *SystemLog) GetEventName() *string {
	if s == nil {
		return nil
	}
	return s.EventName
}

func (s *SystemLog) SetEventName(eventName string) {
	s.EventName = &eventName
}

func (s *SystemLog) GetEventTime() *string {
	if s == nil {
		return nil
	}
	return s.EventTime
}

func (s *SystemLog) SetEventTime(eventTime string) {
	s.EventTime = &eventTime
}

func (s *SystemLog) GetInstanceId() *string {
	if s == nil {
		return nil
	}
	return s.InstanceId
}

func (s *SystemLog) SetInstanceId(instanceId string) {
	s.InstanceId = &instanceId
}

func (s *SystemLog) GetMachineId() *string {
	if s == nil {
		return nil
	}
	return s.MachineId
}

func (s *SystemLog) SetMachineId(machineId string) {
	s.MachineId = &machineId
}

func (s *SystemLog) GetVersion() *string {
	if s == nil {
		return nil
	}
	return s.Version
}

func (s *SystemLog) SetVersion(version string) {
	s.Version = &version
}

func (s *SystemLog) GetResourceCpu() *util.Nullable[int64] {
	if s == nil {
		return nil
	}
	return s.ResourceCpu
}

func (s *SystemLog) SetResourceCpu(resourceCpu util.Nullable[int64]) {
	s.ResourceCpu = &resourceCpu
}

func (s *SystemLog) SetResourceCpuNull() {
	s.ResourceCpu = &util.Nullable[int64]{IsNull: true}
}

func (s *SystemLog) GetResourceMemory() *util.Nullable[int64] {
	if s == nil {
		return nil
	}
	return s.ResourceMemory
}

func (s *SystemLog) SetResourceMemory(resourceMemory util.Nullable[int64]) {
	s.ResourceMemory = &resourceMemory
}

func (s *SystemLog) SetResourceMemoryNull() {
	s.ResourceMemory = &util.Nullable[int64]{IsNull: true}
}

func (s *SystemLog) GetResourceGpuClass() *string {
	if s == nil {
		return nil
	}
	return s.ResourceGpuClass
}

func (s *SystemLog) SetResourceGpuClass(resourceGpuClass string) {
	s.ResourceGpuClass = &resourceGpuClass
}

func (s *SystemLog) GetResourceStorageAmount() *util.Nullable[int64] {
	if s == nil {
		return nil
	}
	return s.ResourceStorageAmount
}

func (s *SystemLog) SetResourceStorageAmount(resourceStorageAmount util.Nullable[int64]) {
	s.ResourceStorageAmount = &resourceStorageAmount
}

func (s *SystemLog) SetResourceStorageAmountNull() {
	s.ResourceStorageAmount = &util.Nullable[int64]{IsNull: true}
}

func (s SystemLog) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SystemLog to string"
	}
	return string(jsonData)
}
