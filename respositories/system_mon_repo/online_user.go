package system_mon_repo

import (
	"cook-book-backEnd/models"
	"gorm.io/gorm"
)

type OnlineUserRepository struct {
	db *gorm.DB
}

func NewOnlineUserRepository(db *gorm.DB) *OnlineUserRepository {
	return &OnlineUserRepository{db: db}
}

func (r *OnlineUserRepository) GetOnlineUsers() ([]models.AdminUser, int64, error) {
	var onlineUsers []models.AdminUser
	var count int64

	err := r.db.Where("login_status =1").Find(&onlineUsers).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return onlineUsers, count, nil
}
