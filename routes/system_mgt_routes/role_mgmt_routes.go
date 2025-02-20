package system_mgt_routes

import (
	"cook-book-backEnd/controllers/system_mgmt_ctrl"
	"cook-book-backEnd/respositories"
	"cook-book-backEnd/respositories/system_mgmt_repo"
	"cook-book-backEnd/services/system_mgmt_srv"
	"github.com/gin-gonic/gin"
)

func SetupRoleMgmtRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	roleMgmtController := system_mgmt_ctrl.NewRoleMgmtController(system_mgmt_srv.NewRoleMgmtService(system_mgmt_repo.NewRoleMgmtRepository(respositories.DB)))
	rg.POST("/admin/role", roleMgmtController.GetRoles)
	rg.POST("/admin/role/update", roleMgmtController.UpdateRole)
	rg.POST("/admin/role/add", roleMgmtController.CreateRole)
	rg.GET("/admin/role/delete", roleMgmtController.DeleteRole)
	rg.POST("/admin/role/role-menu-ids", roleMgmtController.GetRoleMenuListByRoleId)
	rg.POST("/admin/role/save-role-menus", roleMgmtController.SaveRoleMenuPermission)

}
