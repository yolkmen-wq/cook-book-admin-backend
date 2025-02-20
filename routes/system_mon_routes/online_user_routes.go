package system_mon_routes

import (
	"cook-book-backEnd/controllers/system_mon_ctrl"
	"cook-book-backEnd/respositories"
	"cook-book-backEnd/respositories/system_mon_repo"
	"cook-book-backEnd/services/system_mon_srv"
	"github.com/gin-gonic/gin"
)

func SetupOnlineUserRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	onlineUserController := system_mon_ctrl.NewOnlineUserController(system_mon_srv.NewOnlineUserService(system_mon_repo.NewOnlineUserRepository(respositories.DB)))

	rg.POST("/admin/online-user", onlineUserController.GetOnlineUsers)

}
