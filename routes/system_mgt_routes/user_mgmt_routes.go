package system_mgt_routes

import (
	"cook-book-backEnd/controllers/system_mgmt_ctrl"
	"cook-book-backEnd/respositories"
	"cook-book-backEnd/respositories/system_mgmt_repo"
	"cook-book-backEnd/services/system_mgmt_srv"
	"github.com/gin-gonic/gin"
)

func SetupUserMgmtRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	userMgmtController := system_mgmt_ctrl.NewUserMgmtController(system_mgmt_srv.NewUserMgmtService(system_mgmt_repo.NewUserMgmtRepository(respositories.DB)))

	rg.POST("/admin/user", userMgmtController.GetUsers)
	rg.POST("/admin/user/update", userMgmtController.UpdateUser)
	rg.POST("/admin/user/delete", userMgmtController.DeleteUser)
	rg.POST("/admin/list-all-role", userMgmtController.GetRoles)
	rg.POST("/admin/list-role-ids", userMgmtController.GetRolesByIds)
	rg.POST("/admin/user/assignRole", userMgmtController.AssignRolesToUser)
}
