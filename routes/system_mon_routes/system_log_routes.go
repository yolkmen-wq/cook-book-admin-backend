package system_mon_routes

import (
	"cook-book-backEnd/controllers/system_mon_ctrl"
	"cook-book-backEnd/respositories"
	"cook-book-backEnd/respositories/system_mon_repo"
	"cook-book-backEnd/services/system_mon_srv"
	"github.com/gin-gonic/gin"
)

func SetupSystemLogRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	systemLogController := system_mon_ctrl.NewSystemLogController(system_mon_srv.NewSystemLogService(system_mon_repo.NewSystemLogRepository(respositories.DB)))
	rg.POST("/admin/system-logs", systemLogController.GetSystemLogs)
	rg.POST("/admin/system-logs/delete", systemLogController.DeleteSystemLogs)
}
