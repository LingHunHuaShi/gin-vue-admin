import service from '@/utils/request'

// CreateTask 注册任务创建函数
// 参数 task 任务对象
// 返回值 被添加的任务对象 错误信息
export const createTask = (data) => {
  return service({
    url: '/task/createTask',
    method: 'post',
    data
  })
}

// FindTaskByUser 根据用户查找任务
// 参数 user 系统用户对象
// 返回值 查找到的任务对象指针 错误信息
export const findTaskByUser = (data) => {
  return service({
    url: '/task/findTaskByUser',
    method: 'post',
    data
  })
}

// FindTaskByTaskID 根据任务ID查找任务
// 参数 taskID 任务ID
// 返回值 查找到的任务对象指针 错误信息
export const findTaskByTaskID = (data) => {
  return service({
    url: '/task/findTaskByTaskID',
    method: 'post',
    data
  })
}

// DeleteTask 根据任务ID删除任务
// 参数 taskID 任务ID
// 返回值 错误信息
export const deleteTask = (data) => {
  return service({
    url: '/task/deleteTask',
    method: 'delete',
    data
  })
}

// UpdateTaskByTaskID 根据任务ID更新任务信息
// 参数 taskID 任务ID, newTask 新的任务信息
// 返回值 更新后的任务对象指针 错误信息
export const updateTask = (data) => {
  return service({
    url: '/task/updateTask',
    method: 'put',
    data
  })
}

export const queryOngoingTask = (data) => {
  return service({
    url: '/task/queryOngoingTask',
    method: 'get',
    data
  })
}

export const startProcess = (data) => {
  return service({
    url: '/task/startProcess' ,
    method: 'post',
    data
  })
}

export const pauseProcess = (data) => {
  return service({
    url: '/task/pauseProcess',
    method: 'post',
    data
  })
}

export const awakeProcess = (data) => {
  return service({
    url: '/task/awakeProcess',
    method: 'post',
    data
  })
}

export const killProcess = (data) => {
  return service({
    url: '/task/killProcess',
    method: 'post',
    data
  })
}
