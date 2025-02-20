package system_mgmt_repo

import (
	"cook-book-backEnd/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type RoleMgmtRepository struct {
	db *gorm.DB
}

func NewRoleMgmtRepository(db *gorm.DB) *RoleMgmtRepository {
	return &RoleMgmtRepository{db: db}
}

// 获取角色列表
func (lr *RoleMgmtRepository) GetRoleList(searchInfo models.GetRolesRequest) ([]models.Role, int64, error) {
	var roles []models.Role
	var db = lr.db
	var count int64

	// 构建查询条件
	if searchInfo.ID != nil {
		fmt.Println("id", *searchInfo.ID)
		db = db.Joins("LEFT JOIN user_roles ON user_roles.role_id = roles.role_id").Where("user_roles.user_id = ?", *searchInfo.ID).Count(&count)
	}

	if searchInfo.Name != "" {
		fmt.Println("name", searchInfo.Name)
		db = db.Where("role_name like ?", "%"+searchInfo.Name+"%")
	}

	if searchInfo.Code != "" {
		fmt.Println("code", searchInfo.Code)
		db = db.Where("role_code like ?", "%"+searchInfo.Code+"%")
	}

	if searchInfo.Status != nil && searchInfo.Status != "" {
		fmt.Println("status", searchInfo.Status)
		db = db.Where("role_status = ?", searchInfo.Status)
	}

	db = db.Table("roles").Find(&roles)
	fmt.Println("roles", roles)
	if err := db.Error; err != nil {
		fmt.Println("获取角色列表失败", err)
		return nil, 0, err
	}
	return roles, count, nil

}

// 删除角色
func (lr *RoleMgmtRepository) DeleteRole(id int64) error {
	var role models.Role
	if err := lr.db.Where("role_id = ?", id).Delete(&role).Error; err != nil {
		fmt.Println("删除角色失败", err)
		return err
	}
	return nil
}

// 修改角色
func (lr *RoleMgmtRepository) UpdateRole(roleInfo models.Role) error {
	var role models.Role
	if err := lr.db.Where("role_id = ?", roleInfo.ID).First(&role).Error; err != nil {
		fmt.Println("修改角色失败", err)
		return err
	}
	role.Updatetime = time.Now()
	if roleInfo.Name != "" {
		fmt.Println("Name", roleInfo.Name)
		role.Name = roleInfo.Name
	}
	if roleInfo.Code != "" {
		fmt.Println("code", roleInfo.Code)
		role.Code = roleInfo.Code
	}
	if &roleInfo.Status != nil {
		fmt.Println("status", roleInfo.Status)
		role.Status = roleInfo.Status
	}
	if err := lr.db.Save(&role).Error; err != nil {
		fmt.Println("修改角色失败", err)
		return err
	}
	return nil
}

// 新增角色
func (lr *RoleMgmtRepository) CreateRole(roleName, code string) error {
	var role models.Role
	role.Name = roleName
	role.Code = code
	fmt.Println("role", role)
	if err := lr.db.Create(&role).Error; err != nil {
		fmt.Println("新增角色失败", err)
		return err
	}
	return nil
}

// 获取菜单列表
func (lr *RoleMgmtRepository) GetRoleMenuListByRoleId(roleId int64) ([]int64, error) {
	var roleMenus []models.RoleMenu
	var menusId []int64
	fmt.Println("roleId", roleId)
	err := lr.db.Table("role_menus").Where("role_id = ?", roleId).Select("menu_id").Find(&roleMenus).Error
	if err != nil {
		fmt.Println("获取菜单列失败", err)
		return nil, err
	}
	for _, roleMenu := range roleMenus {
		menusId = append(menusId, roleMenu.MenuID)
	}
	fmt.Println("menusId", menusId)

	return menusId, nil
}

// 保存角色菜单权限
func (lr *RoleMgmtRepository) SaveRoleMenuPermission(roleId int64, menuIds []int64) error {
	// 删除原有权限
	lr.db.Where("role_id = ?", roleId).Delete(&models.RoleMenu{})

	// 保存新权限
	for _, menuId := range menuIds {
		rolePermission := models.RoleMenu{
			RoleID: roleId,
			MenuID: menuId,
		}
		if err := lr.db.Create(&rolePermission).Error; err != nil {
			fmt.Println("保存角色菜单权限失败", err)
			return err
		}
	}
	return nil
}
