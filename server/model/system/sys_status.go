package system

type SysStatus struct {
	DeviceModel   string  `json:"deviceModel"`
	CpuUsage      float32 `json:"cpuUSage"`
	NpuUsage      float32 `json:"npuUsage"`
	RamUsage      float32 `json:"ramUsage"`
	StorageUsage  float32 `json:"storageUsage"`
	NumberOfTasks uint    `json:"numberOfTasks"`
}
