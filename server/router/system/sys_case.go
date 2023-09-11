package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CaseRouter struct{}

func InitCaseRouter(Router *gin.RouterGroup) {
	CaseRouter := Router.Group("case").Use(middleware.OperationRecord())
	CaseApi := v1.ApiGroupApp.SystemApiGroup.CaseApi
	{
		CaseRouter.POST("CreateCase", CaseApi.CreateCase)
		CaseRouter.GET("FindCaseByCaseID", CaseApi.FindCaseByCaseID)
		CaseRouter.GET("FindCaseByUUID", CaseApi.FindCaseByUUID)
		CaseRouter.PUT("UpdateCase", CaseApi.UpdateCase)
		CaseRouter.DELETE("DeleteCase", CaseApi.DeleteCase)
	}
}
