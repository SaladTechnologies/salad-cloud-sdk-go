package systemlogs

import (
	"encoding/json"
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
	ResourceCpu *int64 `json:"resource_cpu,omitempty" required:"true" min:"1" max:"16"`
	// The memory amount in MB
	ResourceMemory *int64 `json:"resource_memory,omitempty" required:"true" min:"1024" max:"61440"`
	// The GPU class name
	ResourceGpuClass *string `json:"resource_gpu_class,omitempty" required:"true"`
	// The storage amount in bytes
	ResourceStorageAmount *int64 `json:"resource_storage_amount,omitempty" required:"true" min:"1073741824" max:"53687091200"`
	touched               map[string]bool
}

func (s *SystemLog) GetEventName() *string {
	if s == nil {
		return nil
	}
	return s.EventName
}

func (s *SystemLog) SetEventName(eventName string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["EventName"] = true
	s.EventName = &eventName
}

func (s *SystemLog) SetEventNameNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["EventName"] = true
	s.EventName = nil
}

func (s *SystemLog) GetEventTime() *string {
	if s == nil {
		return nil
	}
	return s.EventTime
}

func (s *SystemLog) SetEventTime(eventTime string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["EventTime"] = true
	s.EventTime = &eventTime
}

func (s *SystemLog) SetEventTimeNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["EventTime"] = true
	s.EventTime = nil
}

func (s *SystemLog) GetInstanceId() *string {
	if s == nil {
		return nil
	}
	return s.InstanceId
}

func (s *SystemLog) SetInstanceId(instanceId string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["InstanceId"] = true
	s.InstanceId = &instanceId
}

func (s *SystemLog) SetInstanceIdNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["InstanceId"] = true
	s.InstanceId = nil
}

func (s *SystemLog) GetMachineId() *string {
	if s == nil {
		return nil
	}
	return s.MachineId
}

func (s *SystemLog) SetMachineId(machineId string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["MachineId"] = true
	s.MachineId = &machineId
}

func (s *SystemLog) SetMachineIdNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["MachineId"] = true
	s.MachineId = nil
}

func (s *SystemLog) GetVersion() *string {
	if s == nil {
		return nil
	}
	return s.Version
}

func (s *SystemLog) SetVersion(version string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Version"] = true
	s.Version = &version
}

func (s *SystemLog) SetVersionNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Version"] = true
	s.Version = nil
}

func (s *SystemLog) GetResourceCpu() *int64 {
	if s == nil {
		return nil
	}
	return s.ResourceCpu
}

func (s *SystemLog) SetResourceCpu(resourceCpu int64) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ResourceCpu"] = true
	s.ResourceCpu = &resourceCpu
}

func (s *SystemLog) SetResourceCpuNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ResourceCpu"] = true
	s.ResourceCpu = nil
}

func (s *SystemLog) GetResourceMemory() *int64 {
	if s == nil {
		return nil
	}
	return s.ResourceMemory
}

func (s *SystemLog) SetResourceMemory(resourceMemory int64) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ResourceMemory"] = true
	s.ResourceMemory = &resourceMemory
}

func (s *SystemLog) SetResourceMemoryNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ResourceMemory"] = true
	s.ResourceMemory = nil
}

func (s *SystemLog) GetResourceGpuClass() *string {
	if s == nil {
		return nil
	}
	return s.ResourceGpuClass
}

func (s *SystemLog) SetResourceGpuClass(resourceGpuClass string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ResourceGpuClass"] = true
	s.ResourceGpuClass = &resourceGpuClass
}

func (s *SystemLog) SetResourceGpuClassNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ResourceGpuClass"] = true
	s.ResourceGpuClass = nil
}

func (s *SystemLog) GetResourceStorageAmount() *int64 {
	if s == nil {
		return nil
	}
	return s.ResourceStorageAmount
}

func (s *SystemLog) SetResourceStorageAmount(resourceStorageAmount int64) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ResourceStorageAmount"] = true
	s.ResourceStorageAmount = &resourceStorageAmount
}

func (s *SystemLog) SetResourceStorageAmountNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ResourceStorageAmount"] = true
	s.ResourceStorageAmount = nil
}

func (s SystemLog) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["EventName"] && s.EventName == nil {
		data["event_name"] = nil
	} else if s.EventName != nil {
		data["event_name"] = s.EventName
	}

	if s.touched["EventTime"] && s.EventTime == nil {
		data["event_time"] = nil
	} else if s.EventTime != nil {
		data["event_time"] = s.EventTime
	}

	if s.touched["InstanceId"] && s.InstanceId == nil {
		data["instance_id"] = nil
	} else if s.InstanceId != nil {
		data["instance_id"] = s.InstanceId
	}

	if s.touched["MachineId"] && s.MachineId == nil {
		data["machine_id"] = nil
	} else if s.MachineId != nil {
		data["machine_id"] = s.MachineId
	}

	if s.touched["Version"] && s.Version == nil {
		data["version"] = nil
	} else if s.Version != nil {
		data["version"] = s.Version
	}

	if s.touched["ResourceCpu"] && s.ResourceCpu == nil {
		data["resource_cpu"] = nil
	} else if s.ResourceCpu != nil {
		data["resource_cpu"] = s.ResourceCpu
	}

	if s.touched["ResourceMemory"] && s.ResourceMemory == nil {
		data["resource_memory"] = nil
	} else if s.ResourceMemory != nil {
		data["resource_memory"] = s.ResourceMemory
	}

	if s.touched["ResourceGpuClass"] && s.ResourceGpuClass == nil {
		data["resource_gpu_class"] = nil
	} else if s.ResourceGpuClass != nil {
		data["resource_gpu_class"] = s.ResourceGpuClass
	}

	if s.touched["ResourceStorageAmount"] && s.ResourceStorageAmount == nil {
		data["resource_storage_amount"] = nil
	} else if s.ResourceStorageAmount != nil {
		data["resource_storage_amount"] = s.ResourceStorageAmount
	}

	return json.Marshal(data)
}

func (s SystemLog) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SystemLog to string"
	}
	return string(jsonData)
}
