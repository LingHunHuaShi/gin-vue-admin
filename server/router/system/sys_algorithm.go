package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlgorithmRouter struct{}

func (s *AlgorithmRouter) InitAlgorithmRouter(Router *gin.RouterGroup) {
	algorithmRouter := Router.Group("algorithm").Use(middleware.OperationRecord())
	algorithmApi := v1.ApiGroupApp.SystemApiGroup.AlgorithmApi
	{
		algorithmRouter.POST("createAlgorithm", algorithmApi.CreateAlgorithm)
	}
}
