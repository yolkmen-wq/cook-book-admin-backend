package models

import "time"

type UserRole struct {
	UserID int64 `gorm:"column:user_id"`
	RoleID int64 `gorm:"column:role_id"`
}

type Role struct {
	ID         int64     `json:"id" gorm:"column:role_id;primary_key"`
	Code       string    `json:"code" gorm:"column:role_code"`
	Name       string    `json:"name" gorm:"column:role_name"`
	Status     int       `json:"status" gorm:"column:role_status;default:1"`
	Createtime time.Time `json:"createTime" gorm:"column:create_at;autoCreateTime"`
	Updatetime time.Time `json:"updateTime" gorm:"column:update_at;autoUpdateTime"`
}

type RolePermission struct {
	RoleID       int64 `gorm:"column:role_id"`
	PermissionID int64 `gorm:"column:permission_id"`
}

type RoleMenu struct {
	ID     int64 `gorm:"column:id;primary_key"`
	MenuID int64 `gorm:"column:menu_id"`
	RoleID int64 `gorm:"column:role_id"`
}

type GetRolesRequest struct {
	ID       *int64      `json:"id"`
	Name     string      `json:"name"`
	Code     string      `json:"code"`
	Status   interface{} `json:"status,omitempty"`
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
}
