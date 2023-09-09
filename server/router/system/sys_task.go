package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskRouter struct{}

func InitTaskRouter(Router *gin.RouterGroup) {
	TaskRouter := Router.Group("task").Use(middleware.OperationRecord())
	TaskApi := v1.ApiGroupApp.SystemApiGroup.TaskApi
	{
		TaskRouter.POST("CreateTask", TaskApi.CreateTask)
		TaskRouter.GET("FindTaskByTaskID", TaskApi.FindTaskByTaskID)
		TaskRouter.GET("FindTaskByUser", TaskApi.FindTaskByUser)
		TaskRouter.GET("UpdateTask", TaskApi.UpdateTask)
		TaskRouter.DELETE("DeleteTask", TaskApi.DeleteTask)
	}
}
