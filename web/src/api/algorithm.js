import service from '@/utils/request'

// @Tags algorithm
// @Summary 获取算法列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getLocalAlgorithmList [get]
export const getLocalAlgorithmList = (data) => {
  return service({
    url: '/api/getLocalAlgorithmList',
    method: 'get',
    data
  })
}
