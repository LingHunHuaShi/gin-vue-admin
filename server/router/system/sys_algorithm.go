package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlgorithmRouter struct{}

func InitAlgorithmRouter(Router *gin.RouterGroup) {
	AlgorithmRouter := Router.Group("algorithm").Use(middleware.OperationRecord())
	{
		//AlgorithmRouter.POST("createAlgorithm", system.AlgorithmService{}.CreateAlgorithm)
	}
}
