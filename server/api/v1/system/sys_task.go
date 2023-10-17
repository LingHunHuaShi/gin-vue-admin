package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TaskApi struct{}

// CreateTask 算法任务创建
// 参数 Task
// 返回值 被添加的Task 错误信息
func (Api *TaskApi) CreateTask(c *gin.Context) {
	var Task system.SysTask
	err := c.ShouldBindJSON(&Task)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = taskService.CreateTask(Task)
	if err != nil {
		global.GVA_LOG.Error("创建任务失败!", zap.Error(err))
		response.FailWithMessage("创建任务失败", c)
		return
	}
	response.OkWithMessage("创建任务成功", c)
}

// FindTaskByUser 根据用户查找任务
// 参数 user 系统用户对象
// 返回值 查找到的任务对象指针 错误信息
func (Api *TaskApi) FindTaskByUser(c *gin.Context) {
	var User system.SysUser
	// var Task *system.SysTask
	err := c.ShouldBindJSON(&User)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Task, err := taskService.FindTaskByUser(User)
	if err != nil {
		global.GVA_LOG.Error("查找失败!", zap.Error(err))
		response.FailWithMessage("查找失败!"+err.Error(), c)
		return
	}
	response.OkWithDetailed(Task, "查找成功!", c)
}

// FindTaskByTaskID 根据任务ID查找任务
// 参数 taskID 任务ID
// 返回值 查找到的任务对象指针 错误信息
func (Api *TaskApi) FindTaskByTaskID(c *gin.Context) {
	var Task system.SysTask
	err := c.ShouldBindJSON(&Task)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	task, err := taskService.FindTaskByTaskID(Task.ID)
	if err != nil {
		global.GVA_LOG.Error("查找失败!", zap.Error(err))
		response.FailWithMessage("查找失败!"+err.Error(), c)
		return
	}
	response.OkWithDetailed(task, "查找成功!", c)
}

// DeleteTask 根据任务ID删除任务
// 参数 taskID 任务ID
// 返回值 错误信息
func (Api *TaskApi) DeleteTask(c *gin.Context) {
	var Task system.SysTask
	err := c.ShouldBindJSON(&Task)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = taskService.DeleteTask(Task.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败!"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功!", c)
}

// UpdateTask  根据任务ID更新任务信息
// 参数 taskID 任务ID, newTask 新的任务信息
// 返回值 更新后的任务对象指针 错误信息
func (Api *TaskApi) UpdateTask(c *gin.Context) {
	var Task system.SysTask
	err := c.ShouldBindJSON(&Task)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	newTask, err := taskService.UpdateTaskByTaskID(Task)
	if err != nil {
		global.GVA_LOG.Error("更新失败2!", zap.Error(err))
		response.FailWithMessage("更新失败2!", c)
		return
	}
	response.OkWithDetailed(newTask, "更新成功!", c)
}

// QueryOngoingTask 查询正在运行的任务
// 参数 None
// 返回值 []Task 错误信息
func (Api *TaskApi) QueryOngoingTask(c *gin.Context) {
	var Container []*system.SysTask
	Container, err := taskService.QueryOngoingTask()
	if err != nil {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(Container, "查询成功", c)
}
