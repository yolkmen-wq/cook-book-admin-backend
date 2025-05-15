package emoji_srv

import (
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/respositories/emoji_repo"
)

type EmojiMgmtService interface {
	GetEmojiList(req *models.GetEmojisRequest) ([]models.Emoji, int64, int, int, error)
	AddEmoji(emoji *models.Emoji) error
	UpdateEmoji(emoji *models.Emoji) error
	DeleteEmoji(id int64) error
}

type emojiMgmtService struct {
	repo emoji_repo.EmojiMgmtRepository
}

func NewEmojiMgmtService(repo *emoji_repo.EmojiMgmtRepository) EmojiMgmtService {
	return &emojiMgmtService{
		repo: *repo,
	}
}

// GetAllEmoji returns all emojis with pagination
func (ems *emojiMgmtService) GetEmojiList(req *models.GetEmojisRequest) ([]models.Emoji, int64, int, int, error) {
	return ems.repo.GetEmojiList(req)
}

// AddEmoji adds a new emoji
func (ems *emojiMgmtService) AddEmoji(emoji *models.Emoji) error {
	return ems.repo.AddEmoji(emoji)
}

// UpdateEmoji updates an existing emoji
func (ems *emojiMgmtService) UpdateEmoji(emoji *models.Emoji) error {
	return ems.repo.UpdateEmoji(emoji)
}

// DeleteEmoji deletes an emoji by id
func (ems *emojiMgmtService) DeleteEmoji(id int64) error {
	return ems.repo.DeleteEmoji(id)
}
