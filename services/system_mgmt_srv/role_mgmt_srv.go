package system_mgmt_srv

import (
	"cook-book-backEnd/models"
	"cook-book-backEnd/respositories/system_mgmt_repo"
)

type RoleMgmtService interface {
	GetRoleList(searchInfo models.GetRolesRequest) ([]models.Role, int64, error)
	UpdateRole(roleInfo models.Role) error
	CreateRole(roleName, code string) error
	DeleteRole(id int64) error
	GetRoleMenuListByRoleId(roleId int64) ([]int64, error)
	SaveRoleMenuPermission(roleId int64, menuIds []int64) error
}

type roleMgmtService struct {
	roleMgmtRepo system_mgmt_repo.RoleMgmtRepository
}

func NewRoleMgmtService(roleMgmtRepo *system_mgmt_repo.RoleMgmtRepository) RoleMgmtService {
	return &roleMgmtService{
		roleMgmtRepo: *roleMgmtRepo,
	}
}

// GetRoleList get role list
func (s *roleMgmtService) GetRoleList(searchInfo models.GetRolesRequest) ([]models.Role, int64, error) {
	return s.roleMgmtRepo.GetRoleList(searchInfo)
}

// UpdateRole update role
func (s *roleMgmtService) UpdateRole(roleInfo models.Role) error {
	return s.roleMgmtRepo.UpdateRole(roleInfo)
}

// CreateRole create role
func (s *roleMgmtService) CreateRole(roleName, code string) error {
	return s.roleMgmtRepo.CreateRole(roleName, code)
}

// DeleteRole delete role
func (s *roleMgmtService) DeleteRole(id int64) error {
	return s.roleMgmtRepo.DeleteRole(id)
}

// GetMenuListByRoleId get menu list by role id
func (s *roleMgmtService) GetRoleMenuListByRoleId(roleId int64) ([]int64, error) {
	return s.roleMgmtRepo.GetRoleMenuListByRoleId(roleId)
}

// SaveRoleMenuPermission save role menu permission
func (s *roleMgmtService) SaveRoleMenuPermission(roleId int64, menuIds []int64) error {
	return s.roleMgmtRepo.SaveRoleMenuPermission(roleId, menuIds)
}
