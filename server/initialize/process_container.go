package initialize

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

var ProcessManager *system.SysProcessContainer

func InitProcessManageContainer() *system.SysProcessContainer {
	return &system.SysProcessContainer{}
}
