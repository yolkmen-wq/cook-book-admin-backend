package system_mgmt_ctrl

import (
	"cook-book-backEnd/config"
	"cook-book-backEnd/models"
	"cook-book-backEnd/services/system_mgmt_srv"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MenuMgmtController struct {
	menuMgmtService system_mgmt_srv.MenuMgmtService
}

func NewMenuMgmtController(menuMgmtService system_mgmt_srv.MenuMgmtService) *MenuMgmtController {
	return &MenuMgmtController{menuMgmtService}
}

func (m *MenuMgmtController) GetMenus(c *gin.Context) {

	var request models.GetMenuRequest
	if err := c.ShouldBind(&request); err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	menus, total, err := m.menuMgmtService.GetMenuList(request)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", map[string]interface{}{
		"list":  menus,
		"total": total,
	})
	c.JSONP(http.StatusOK, response)
}

func (m *MenuMgmtController) GetMenuDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		errResponse := config.NewResponse(http.StatusBadRequest, false, "参数错误", nil)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	menu, err := m.menuMgmtService.GetMenuDetail(parsedId)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", menu)
	c.JSONP(http.StatusOK, response)
}

func (m *MenuMgmtController) AddMenu(c *gin.Context) {
	var req models.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	err := m.menuMgmtService.CreateMenu(req)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "添加成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (m *MenuMgmtController) UpdateMenu(c *gin.Context) {
	var req models.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	err := m.menuMgmtService.UpdateMenu(req)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "更新成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (m *MenuMgmtController) DeleteMenu(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		errResponse := config.NewResponse(http.StatusBadRequest, false, "参数错误", nil)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	err = m.menuMgmtService.DeleteMenu(parsedId)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "删除成功", nil)
	c.JSONP(http.StatusOK, response)
}
