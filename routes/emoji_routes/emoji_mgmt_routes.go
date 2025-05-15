package emoji_routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers/emoji_ctrl"
	"cook-book-admin-backend/respositories/emoji_repo"
	"cook-book-admin-backend/services/emoji_srv"
	"github.com/gin-gonic/gin"
)

func SetupEmojiMgmtRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	articleMgmtController := emoji_ctrl.NewEmojiMgmtController(emoji_srv.NewEmojiMgmtService(emoji_repo.NewEmojiMgmtRepository(config.DB)))

	r.POST("/admin/emoji", articleMgmtController.GetEmojiList)
	r.POST("/admin/emoji/create", articleMgmtController.CreateEmoji)
	r.POST("/admin/emoji/update", articleMgmtController.UpdateEmoji)
	r.GET("/admin/emoji/delete/:id", articleMgmtController.DeleteEmoji)
}
