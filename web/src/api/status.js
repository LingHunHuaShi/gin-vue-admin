import service from '@/utils/request'

// uploadSystemStatus 向云端系统报告本地盒子当前的状态
// @param
// DeviceModel   string  `json:"deviceModel"`
// CpuUsage      float32 `json:"cpuUSage"`
// NpuUsage      float32 `json:"npuUsage"`
// RamUsage      float32 `json:"ramUsage"`
// StorageUsage  float32 `json:"storageUsage"`
// NumberOfTasks uint    `json:"numberOfTasks"`
export const uploadSystemStatus = (data) => {
  return service({
    url: '/status/uploadSystemStatus',
    method: 'post',
    data
  })
}
