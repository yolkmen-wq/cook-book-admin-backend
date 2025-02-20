package system_mgt_routes

import (
	"cook-book-backEnd/controllers/system_mgmt_ctrl"
	"cook-book-backEnd/respositories"
	"cook-book-backEnd/respositories/system_mgmt_repo"
	"cook-book-backEnd/services/system_mgmt_srv"
	"github.com/gin-gonic/gin"
)

func SetupMenuMgmtRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	menuMgmtController := system_mgmt_ctrl.NewMenuMgmtController(system_mgmt_srv.NewMenuMgmtService(system_mgmt_repo.NewMenuMgmtRepository(respositories.DB)))
	rg.POST("/admin/menu", menuMgmtController.GetMenus)
	rg.GET("admin/menu/detail", menuMgmtController.GetMenuDetail)
	rg.POST("admin/menu/update", menuMgmtController.UpdateMenu)
	rg.POST("admin/menu/create", menuMgmtController.AddMenu)
	rg.GET("admin/menu/delete", menuMgmtController.DeleteMenu)
}
