package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CaseRouter struct{}

func InitCaseRouter(Router *gin.RouterGroup) {
	CaseRouter := Router.Group("case").Use(middleware.OperationRecord())
	{

	}
}
