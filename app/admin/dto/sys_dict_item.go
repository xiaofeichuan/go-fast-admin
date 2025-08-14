package dto

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common/dto/request"
	"time"
)

// SysDictItemQuery 查询
type SysDictItemQuery struct {
	request.PageQuery
}

// SysDictItemVo 输出对象
type SysDictItemVo struct {
	Id            int64                 `json:"id"`            //编号
	DictId        int64                 `json:"dictId"`        //字典id
	DictItemName  string                `json:"dictItemName"`  //选项名称
	DictItemValue string                `json:"dictItemValue"` //选项值
	Status        consts.DictItemStatus `json:"status"`        //配置状态（0正常 1停用）
	CreatedAt     time.Time             `json:"createdAt"`     //创建时间
	UpdatedAt     time.Time             `json:"updatedAt"`     //更新时间
	Remark        string                `json:"remark"`        //备注
}

// SysDictItemAddDto 添加
type SysDictItemAddDto struct {
	DictId        int64                 `json:"dictId"`        //字典id
	DictItemName  string                `json:"dictItemName"`  //选项名称
	DictItemValue string                `json:"dictItemValue"` //选项值
	Status        consts.DictItemStatus `json:"status"`        //配置状态（0正常 1停用）
	Remark        string                `json:"remark"`        //备注
}

// SysDictItemUpdateDto 更新
type SysDictItemUpdateDto struct {
	Id int64 `json:"id"` //编号
	SysDictItemAddDto
}
