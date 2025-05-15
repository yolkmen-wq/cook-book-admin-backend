package emoji_ctrl

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/services/emoji_srv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmojiMgmtController struct {
	emojiService emoji_srv.EmojiMgmtService
}

func NewEmojiMgmtController(emojiService emoji_srv.EmojiMgmtService) *EmojiMgmtController {
	return &EmojiMgmtController{
		emojiService: emojiService,
	}
}

func (emc *EmojiMgmtController) CreateEmoji(c *gin.Context) {
	var req models.Emoji
	if err := c.ShouldBindJSON(&req); err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, "请求参数错误", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := emc.emojiService.AddEmoji(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, "请求参数错误", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "创建成功", nil)
	c.JSON(http.StatusOK, response)
}

func (emc *EmojiMgmtController) GetEmojiList(c *gin.Context) {
	var req models.GetEmojisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(45, err)
		response := config.NewResponse(http.StatusBadRequest, false, "请求参数错误", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	articleList, total, pageSize, pageNum, err := emc.emojiService.GetEmojiList(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", config.ListResponse{List: articleList, Total: total, CurrentPage: pageNum, PageSize: pageSize})
	c.JSONP(http.StatusOK, response)
}

func (emc *EmojiMgmtController) UpdateEmoji(c *gin.Context) {
	var req models.Emoji
	if err := c.ShouldBindJSON(&req); err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, "请求参数错误", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := emc.emojiService.UpdateEmoji(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "更新成功", nil)
	c.JSON(http.StatusOK, response)
}

func (emc *EmojiMgmtController) DeleteEmoji(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, "请求参数错误", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = emc.emojiService.DeleteEmoji(idInt)

	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "删除成功", nil)
	c.JSON(http.StatusOK, response)
}
