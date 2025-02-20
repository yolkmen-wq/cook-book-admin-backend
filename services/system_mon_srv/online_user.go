package system_mon_srv

import (
	"cook-book-backEnd/models"
	"cook-book-backEnd/respositories/system_mon_repo"
)

type OnlineUserService interface {
	GetOnlineUsers() ([]models.AdminUser, int64, error)
}

type onlineUserService struct {
	repo system_mon_repo.OnlineUserRepository
}

func NewOnlineUserService(repo *system_mon_repo.OnlineUserRepository) OnlineUserService {
	return &onlineUserService{
		repo: *repo,
	}
}

func (s *onlineUserService) GetOnlineUsers() ([]models.AdminUser, int64, error) {
	return s.repo.GetOnlineUsers()
}
