package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskRouter struct{}

func (s *TaskRouter) InitTaskRouter(Router *gin.RouterGroup) {
	TaskRouter := Router.Group("task").Use(middleware.OperationRecord())
	TaskApi := v1.ApiGroupApp.SystemApiGroup.TaskApi
	{
		TaskRouter.POST("createTask", TaskApi.CreateTask)
		TaskRouter.POST("findTaskByTaskID", TaskApi.FindTaskByTaskID)
		TaskRouter.POST("findTaskByUser", TaskApi.FindTaskByUser)
		TaskRouter.GET("queryOngoingTask", TaskApi.QueryOngoingTask)
		TaskRouter.PUT("updateTask", TaskApi.UpdateTask)
		TaskRouter.DELETE("deleteTask", TaskApi.DeleteTask)
		TaskRouter.POST("startProcess", TaskApi.StartTask)
		TaskRouter.POST("pauseProcess", TaskApi.PauseTask)
		TaskRouter.POST("awakeProcess", TaskApi.AwakeTask)
		TaskRouter.POST("killProcess", TaskApi.KillTask)
	}
}
