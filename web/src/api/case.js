import service from '@/utils/request'

// CreateCase 注册案例创建函数
// 参数 caseObj 案例对象
// 返回值 被添加的案例对象 错误信息
export const createCase = (data) => {
  return service({
    url: '/case/createCase',
    method: 'post',
    data
  })
}

// FindCaseByUUID 根据发起者UUID查找案例
// 参数 uuid 发起者UUID
// 返回值 查找到的案例对象指针 错误信息
export const findCaseByUser = (data) => {
  return service({
    url: '/case/findCaseByUser',
    method: 'post',
    data
  })
}

// FindCaseByCaseID 根据案例ID查找案例
// 参数 caseID 案例ID
// 返回值 查找到的案例对象指针 错误信息
export const findCaseByCaseID = (data) => {
  return service({
    url: '/case/findCaseByCaseID',
    method: 'post',
    data
  })
}

// DeleteCase 根据案例ID删除案例
// 参数 caseID 案例ID
// 返回值 错误信息
export const deleteCase = (data) => {
  return service({
    url: '/case/deleteCase',
    method: 'delete',
    data
  })
}

// UpdateCaseByCaseID 根据案例ID更新案例信息
// 参数 caseID 案例ID, newCase 新的案例信息
// 返回值 更新后的案例对象指针 错误信息
export const updateCase = (data) => {
  return service({
    url: '/case/updateCase',
    method: 'put',
    data
  })
}

export const queryAllCases = (data) => {
  return service({
    url: '/case/queryAllCases',
    method: 'get',
    data
  })
}
