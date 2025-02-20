package system_mgmt_srv

import (
	"cook-book-backEnd/models"
	"cook-book-backEnd/respositories/system_mgmt_repo"
)

type UserMgmtService interface {
	GetAdminUserList(page, size int, username string, nickname string, status interface{}) ([]models.AdminUser, int64, error)
	DeleteUser(id int64) error
	AddUser(info models.AdminUser) error
	UpdateUser(info models.AdminUser) error
	GetRoleList(name string, id int64) ([]models.Role, error)
	UpdateUserRoles(id int64, roleIds []int64) error
}

type userMgmtService struct {
	userMgmtRepo system_mgmt_repo.UserMgmtRepository
}

func NewUserMgmtService(userMgmtRepo *system_mgmt_repo.UserMgmtRepository) UserMgmtService {
	return &userMgmtService{userMgmtRepo: *userMgmtRepo}
}

// GetAdminUserList function is used to get admin user list
func (ums *userMgmtService) GetAdminUserList(page, size int, username string, nickname string, status interface{}) ([]models.AdminUser, int64, error) {
	return ums.userMgmtRepo.GetAdminUserList(page, size, username, nickname, status)
}

// DeleteUser function is used to delete user
func (ums *userMgmtService) DeleteUser(id int64) error {
	return ums.userMgmtRepo.DeleteUser(id)
}

// AddUser function is used to add user
func (ums *userMgmtService) AddUser(info models.AdminUser) error {
	return ums.userMgmtRepo.AddUser(info)
}

// UpdateUserStatus function is used to update user status
func (ums *userMgmtService) UpdateUser(info models.AdminUser) error {
	return ums.userMgmtRepo.UpdateUser(info)
}

// GetRoleList function is used to get role list
func (ums *userMgmtService) GetRoleList(name string, id int64) ([]models.Role, error) {
	return ums.userMgmtRepo.GetUserRoleList(name, id)
}

// UpdateUserRoles function is used to update user roles
func (ums *userMgmtService) UpdateUserRoles(id int64, roleIds []int64) error {
	return ums.userMgmtRepo.UpdateUserRoles(id, roleIds)
}
