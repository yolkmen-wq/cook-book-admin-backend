package emoji_repo

import (
	"cook-book-admin-backend/models"
	"fmt"
	"gorm.io/gorm"
)

type EmojiMgmtRepository struct {
	db *gorm.DB
}

func NewEmojiMgmtRepository(db *gorm.DB) *EmojiMgmtRepository {
	return &EmojiMgmtRepository{db: db}
}

// 获取表情包

func (emp *EmojiMgmtRepository) GetEmojiList(req *models.GetEmojisRequest) ([]models.Emoji, int64, int, int, error) {
	var emojis []models.Emoji
	var total int64

	// 构建查询条件
	db := emp.db.Table("emojis")

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Status != nil {
		db = db.Where("status = ?", req.Status)
	}

	if len(req.CreatedTime) == 2 {
		fmt.Println(req.CreatedTime[0], req.CreatedTime[1])
		db = db.Where("created_at BETWEEN ? AND ?", req.CreatedTime[0], req.CreatedTime[1])
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("获取文章列表总数失败", err)
		return nil, 0, 0, 0, err
	}

	// 分页
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	db = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)

	err := db.Find(&emojis).Error
	if err != nil {
		return nil, 0, 0, 0, err
	}
	return emojis, total, req.PageSize, req.PageNum, nil
}

// 新增表情包
func (emp *EmojiMgmtRepository) AddEmoji(emoji *models.Emoji) error {
	return emp.db.Create(emoji).Error
}

// 更新表情包
func (emp *EmojiMgmtRepository) UpdateEmoji(emoji *models.Emoji) error {
	err := emp.db.Model(&models.Emoji{}).Where("id = ?", emoji.ID).Updates(emoji).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除表情包
func (emp *EmojiMgmtRepository) DeleteEmoji(id int64) error {
	err := emp.db.Delete(&models.Article{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
