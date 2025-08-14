package dto

import (
	"go-fast-admin/server/common/dto/request"
	"time"
)

// SysGenTableQuery 查询
type SysGenTableQuery struct {
	TableName string `json:"tableName" form:"tableName"` //表名称
	request.PageQuery
}

type SysGenTableVo struct {
	Id               int64                 `json:"id" `
	TableName        string                `json:"tableName"`             //表名称
	TableDescription string                `json:"tableDescription"`      //表描述
	ModelName        string                `json:"modelName"`             //实体名称
	BusinessName     string                `json:"businessName"`          //业务名称
	ModuleName       string                `json:"moduleName"`            //模块名称
	ModuleAlias      string                `json:"moduleAlias"  gorm:"-"` //模块别名
	MenuParentId     int64                 `json:"menuParentId"`          //上级菜单
	CreatedAt        time.Time             `json:"createdAt"`             //创建时间
	UpdatedAt        time.Time             `json:"updatedAt"`             //更新时间
	ColumnList       []SysGenTableColumnVo `json:"columnList" gorm:"-"`   //字段列表
}

// SysGenTableAddDto 添加
type SysGenTableAddDto struct {
	TableName        string `json:"tableName"`        //表名称
	TableDescription string `json:"tableDescription"` //表描述
	ModelName        string `json:"modelName"`        //实体名称
	BusinessName     string `json:"businessName"`     //业务名称
	ModuleName       string `json:"moduleName"`       //模块名称
	MenuParentId     int64  `json:"menuParentId"`     //上级菜单
}

// SysGenTableUpdateDto 更新
type SysGenTableUpdateDto struct {
	Id int64 `json:"id"` //编号
	SysGenTableAddDto
}

type TableInfoVo struct {
	TableName        string `json:"tableName"`        //表名称
	TableDescription string `json:"tableDescription"` //表描述
	ModelName        string `json:"modelName"`        //实体名称
	BusinessName     string `json:"businessName"`     //业务名称
	ModuleName       string `json:"moduleName"`       //模块名称
}

type PreviewCodeVo struct {
	FileName    string `json:"fileName"`    //文件名称
	FileContent string `json:"fileContent"` //文件内容
}
