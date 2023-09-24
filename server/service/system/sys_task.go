package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

/*
type SysTask struct {
	global.GVA_MODEL
	TaskID       uint      `json:"taskID" gorm:"default:0;primaryKey;autoIncrement;index;comment:任务ID"` // 任务ID
	UUID         uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                          // 用户UUID
	VideoSource  string    `json:"videoSource" gorm:"comment:视频源"`                                      // 视频源
	CreationDate time.Time `json:"creationDate" gorm:"comment:创建时间"`
	Resolution   int       `json:"resolution" gorm:"comment:分辨率等级"`
	AlgorithmID  uint      `json:"algorithmId" gorm:"comment:任务容器算法ID"`
	Intensity    int       `json:"intensity" gorm:"comment:任务粒度"`
	Status       int       `json:"status" gorm:"comment:任务状态"`
}
*/

type TaskService struct{}

// CreateTask 注册任务创建函数
// 参数 task 任务对象
// 返回值 被添加的任务对象 错误信息
func (t *TaskService) CreateTask(task system.SysTask) (taskInter system.SysTask, err error) {
	err = global.GVA_DB.Create(&task).Error
	if err != nil {
		return task, errors.New("创建任务失败")
	}
	return task, err
}

// FindTaskByUser 根据用户查找任务
// 参数 user 系统用户对象
// 返回值 查找到的任务对象指针 错误信息
func (t *TaskService) FindTaskByUser(user system.SysUser) (taskInter *system.SysTask, err error) {
	var task system.SysTask
	uuid := user.UUID
	// 查找要更新的算法
	err = global.GVA_DB.Where("UUID = ?", uuid).First(&task).Error
	if err != nil {
		return &task, errors.New("任务不存在")
	}
	return &task, err
}

// FindTaskByTaskID 根据任务ID查找任务
// 参数 taskID 任务ID
// 返回值 查找到的任务对象指针 错误信息
func (t *TaskService) FindTaskByTaskID(taskID uint) (taskInter *system.SysTask, err error) {
	var task system.SysTask
	err = global.GVA_DB.Where("ID = ?", taskID).First(&task).Error
	if err != nil {
		return &task, errors.New("任务不存在")
	}
	return &task, err
}

// DeleteTask 根据任务ID删除任务
// 参数 taskID 任务ID
// 返回值 错误信息
func (t *TaskService) DeleteTask(taskID uint) (err error) {
	var task system.SysTask
	err = global.GVA_DB.Where("ID = ?", taskID).Delete(&task).Error
	if err != nil {
		return errors.New("任务不存在")
	}
	return err
}

// UpdateTaskByTaskID 根据任务ID更新任务信息
// 参数 taskID 任务ID, newTask 新的任务信息
// 返回值 更新后的任务对象指针 错误信息
func (t *TaskService) UpdateTaskByTaskID(taskID uint, newTask system.SysTask) (updateTask *system.SysTask, err error) {
	var task system.SysTask

	// 查询任务是否存在
	err = global.GVA_DB.Where("ID = ?", taskID).First(&task).Error
	if err != nil {
		return &task, errors.New("任务不存在")
	}

	// 更新任务信息
	err = global.GVA_DB.Model(&task).Updates(newTask).Error
	if err != nil {
		return &task, errors.New("更新失败")
	}
	return &newTask, err
}

// QueryOngoingTask
// 参数 None
// 返回值 []Task 错误信息
func (t *TaskService) QueryOngoingTask() (Container []*system.SysTask, err error) {
	err = global.GVA_DB.Where("Status = ?", 0).Find(&Container).Error
	if err != nil {
		return nil, errors.New("查询失败")
	}
	return Container, err
}
