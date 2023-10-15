package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CaseRouter struct{}

func (s *CaseRouter) InitCaseRouter(Router *gin.RouterGroup) {
	CaseRouter := Router.Group("case").Use(middleware.OperationRecord())
	CaseApi := v1.ApiGroupApp.SystemApiGroup.CaseApi
	{
		CaseRouter.POST("createCase", CaseApi.CreateCase)
		CaseRouter.POST("findCaseByCaseID", CaseApi.FindCaseByCaseID)
		CaseRouter.POST("findCaseByUUID", CaseApi.FindCaseByUUID)
		CaseRouter.GET("queryAllCases", CaseApi.QueryAllCases)
		CaseRouter.PUT("updateCase", CaseApi.UpdateCase)
		CaseRouter.DELETE("deleteCase", CaseApi.DeleteCase)
	}
}
