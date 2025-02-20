package models

type Menu struct {
	// 菜单类型（0代表菜单、1代表iframe、2代表外链、3代表按钮）
	ID              int64  `json:"id" gorm:"column:id;primary_key"`
	MenuType        int    `json:"menuType" gorm:"column:menu_type"`
	ParentId        int64  `json:"parentId" gorm:"column:parent_id"`
	Title           string `json:"title" gorm:"column:title"`
	Name            string `json:"name" gorm:"column:name"`
	Path            string `json:"path" gorm:"column:path"`
	Component       string `json:"component" gorm:"column:component"`
	Rank            int    `json:"rank" gorm:"column:rank"`
	Redirect        string `json:"redirect" gorm:"column:redirect"`
	Icon            string `json:"icon" gorm:"column:icon"`
	ExtraIcon       string `json:"extraIcon" gorm:"column:extra_icon"`
	EnterTransition string `json:"enterTransition" gorm:"column:enter_transition"`
	LeaveTransition string `json:"leaveTransition" gorm:"column:leave_transition"`
	ActivePath      string `json:"activePath" gorm:"column:active_path"`
	Auths           string `json:"auths" gorm:"column:auths"`
	FrameSrc        string `json:"frameSrc" gorm:"column:frame_src"`
	FrameLoading    bool   `json:"frameLoading" gorm:"column:frame_loading"`
	KeepAlive       bool   `json:"keepAlive" gorm:"column:keep_alive;default:0"`
	HiddenTag       bool   `json:"hiddenTag" gorm:"column:hidden_tag;default:0"`
	FixedTag        bool   `json:"fixedTag" gorm:"column:fixed_tag;default:0"`
	ShowLink        bool   `json:"showLink" gorm:"column:show_link;default:0"`
	ShowParent      bool   `json:"showParent" gorm:"column:show_parent;default:0"`
}

type GetMenuRequest struct {
	Name     string `form:"name"`
	PageNum  int    `form:"page_num"`
	PageSize int    `form:"page_size"`
}
