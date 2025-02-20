package system_mon_ctrl

import (
	"cook-book-backEnd/config"
	"cook-book-backEnd/services/system_mon_srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OnlineUserController struct {
	onlineUserService system_mon_srv.OnlineUserService
}

func NewOnlineUserController(onlineUserService system_mon_srv.OnlineUserService) *OnlineUserController {
	return &OnlineUserController{
		onlineUserService: onlineUserService,
	}
}

func (m *OnlineUserController) GetOnlineUsers(c *gin.Context) {

	users, total, err := m.onlineUserService.GetOnlineUsers()
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", map[string]interface{}{
		"list":  users,
		"total": total,
	})
	c.JSONP(http.StatusOK, response)
}
