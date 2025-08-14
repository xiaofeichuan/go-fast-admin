package dto

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common/dto/request"
	"time"
)

// SysMenuQuery 查询
type SysMenuQuery struct {
	request.PageQuery
	MenuName string `json:"menuName" form:"menuName"` //菜单名称
}

// SysMenuVo 输出对象
type SysMenuVo struct {
	Id         int64             `json:"id"`                //编号
	ParentId   int64             `json:"parentId"`          //父菜单ID
	MenuName   string            `json:"menuName"`          //菜单名称
	MenuType   consts.MenuType   `json:"menuType"`          //菜单类型（0菜单，1按钮）
	Path       string            `json:"path"`              //路由地址
	Link       string            `json:"link"`              //外部地址
	IsIframe   bool              `json:"isIframe"`          //是否内嵌
	IsHide     bool              `json:"isHide"`            //是否隐藏
	IsCache    bool              `json:"isCache"`           //是否缓存
	IsAffix    bool              `json:"isAffix"`           //是否固定tagsView
	Component  string            `json:"component"`         //组件路径
	Icon       string            `json:"icon"`              //菜单图标
	Permission string            `json:"permission"`        //按钮权限
	Status     consts.MenuStatus `json:"status"`            //菜单状态（0正常 1停用）
	Sort       int               `json:"sort"`              //排序
	Remark     string            `json:"remark"`            //备注
	CreatedAt  time.Time         `json:"createdAt"`         //创建时间
	UpdatedAt  time.Time         `json:"updatedAt"`         //更新时间
	Children   []SysMenuVo       `json:"children" gorm:"-"` //子菜单
}

type SysMenuAddDto struct {
	ParentId   int64             `json:"parentId"`   //父菜单ID
	MenuName   string            `json:"menuName"`   //菜单名称
	MenuType   consts.MenuType   `json:"menuType"`   //菜单类型（0菜单，1按钮）
	Path       string            `json:"path"`       //路由地址
	Link       string            `json:"link"`       //外部地址
	IsIframe   bool              `json:"isIframe"`   //是否内嵌
	IsHide     bool              `json:"isHide"`     //是否隐藏
	IsCache    bool              `json:"isCache"`    //是否缓存
	IsAffix    bool              `json:"isAffix"`    //是否固定tagsView
	Component  string            `json:"component"`  //组件路径
	Icon       string            `json:"icon"`       //菜单图标
	Permission string            `json:"permission"` //按钮权限
	Status     consts.MenuStatus `json:"status"`     //菜单状态（0正常 1停用）
	Sort       int               `json:"sort"`       //排序
	Remark     string            `json:"remark"`     //备注
}

// SysMenuUpdateDto 更新
type SysMenuUpdateDto struct {
	Id int64 `json:"id"` //编号
	SysMenuAddDto
}
