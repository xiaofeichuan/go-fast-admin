package dto

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common/dto/request"
	"time"
)

// SysConfigQuery 查询
type SysConfigQuery struct {
	request.PageQuery
	ConfigName string `json:"configName" form:"configName"` //配置名称
}

// SysConfigVo 输出对象
type SysConfigVo struct {
	Id          int64               `json:"id"`          //配置id
	ConfigName  string              `json:"configName"`  //配置名称
	ConfigKey   string              `json:"configKey"`   //配置键名
	ConfigValue string              `json:"configValue"` //配置键值
	Status      consts.ConfigStatus `json:"status"`      //配置状态（0正常 1停用）
	CreatedAt   time.Time           `json:"createdAt"`   //创建时间
	UpdatedAt   time.Time           `json:"updatedAt"`   //更新时间
	Remark      string              `json:"remark"`      //备注
}

// SysConfigAddDto 添加
type SysConfigAddDto struct {
	ConfigName  string              `json:"configName"`  //配置名称
	ConfigKey   string              `json:"configKey"`   //配置键名
	ConfigValue string              `json:"configValue"` //配置键值
	Status      consts.ConfigStatus `json:"status"`      //配置状态（0正常 1停用）
	Remark      string              `json:"remark"`      //备注
}

// SysConfigUpdateDto 更新
type SysConfigUpdateDto struct {
	Id int64 `json:"id"` //编号
	SysConfigAddDto
}
