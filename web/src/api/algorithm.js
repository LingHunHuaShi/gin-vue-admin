import service from '@/utils/request'

// CreateAlgorithm 算法注册
// 参数 algorithm
// 返回值 被添加的算法 错误信息
export const createAlgorithm = (data) => {
  return service({
    url: '/algorithm/createAlgorithm',
    method: 'post',
    data
  })
}

// DeleteAlgorithm 根据算法ID删除算法及其相关权限信息
// @param id uint 算法ID
// @return err error 删除操作的错误，如果删除失败则返回错误信息，否则为nil
export const deleteAlgorithm = (data) => {
  return service({
    url: '/algorithm/deleteAlgorithm',
    method: 'delete',
    data
  })
}

// FindAlgorithmById FindAlgorithmByID 根据算法ID查找算法信息
// @param algorithmID uint 要查找的算法ID
// @return algorithmInter *system.SysAlgorithm 查找到的算法信息，如果找到则为非空指针，否则为nil
// @return err error 查找操作的错误，如果算法不存在或查找失败则返回错误信息，否则为nil
export const findAlgorithmById = (data) => {
  return service({
    url: '/algorithm/findAlgorithmById',
    method: 'get',
    data
  })
}

// UpdateAlgorithm 更新算法信息
// @param algorithmID uint 要更新的算法ID
// @param newAlgorithm system.SysAlgorithm 新的算法信息
// @return updatedAlgorithmInter *system.SysAlgorithm 更新后的算法信息，如果更新成功则为非空指针，否则为nil
// @return err error 更新操作的错误，如果更新失败则返回错误信息，否则为nil
export const updateAlgorithm = (data) => {
  return service({
    url: '/algorithm/updateAlgorithm',
    method: 'put',
    data
  })
}

// QueryAllAlgorithm  查询全部算法信息
// @return algorithms []system.SysAlgorithm 查找到的全部算法信息，如果没有算法则返回空切片
// @return err error 查找操作的错误，如果查找失败则返回错误信息，否则为nil
export const queryAllAlgorithm = (data) => {
  return service({
    url: '/algorithm/queryAllAlgorithm',
    method: 'get',
  })
}
