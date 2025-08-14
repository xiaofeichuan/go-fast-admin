package dto

import (
	"go-fast-admin/server/common/dto/request"
	"time"
)

// SysFileQuery 查询
type SysFileQuery struct {
	request.PageQuery
}

// SysFileAddDto 创建
type SysFileAddDto struct {
	StorageType string `json:"storageType"` //存储类型
	BucketName  string `json:"bucketName"`  //存储桶
	FileName    string `json:"fileName"`    //文件名（如：1.png）
	FileExt     string `json:"fileExt"`     //文件扩展名（如： png）
	MimeType    string `json:"mimeType"`    //文件 MIME 类型（防止伪装文件上传）
	FileSize    int64  `json:"fileSize"`    //文件大小（单位：字节）
	FileUrl     string `json:"fileUrl"`     //文件地址
	IpAddress   string `json:"ipAddress"`   //上传IP
	UserId      int64  `json:"userId"`      //用户Id
}

// SysFileUpdateDto 更新
type SysFileUpdateDto struct {
	Id int64 `json:"id"` //编号
	SysFileAddDto
}

// SysFileVo 输出对象
type SysFileVo struct {
	Id          int64     `json:"id"`          //文件Id
	StorageType string    `json:"storageType"` //存储类型
	BucketName  string    `json:"bucketName"`  //存储桶
	FileName    string    `json:"fileName"`    //文件名（如：1.png）
	FileExt     string    `json:"fileExt"`     //文件扩展名（如： png）
	MimeType    string    `json:"mimeType"`    //文件 MIME 类型（防止伪装文件上传）
	FileSize    int64     `json:"fileSize"`    //文件大小（单位：字节）
	FileUrl     string    `json:"fileUrl"`     //文件地址
	IpAddress   string    `json:"ipAddress"`   //上传IP
	UserId      int64     `json:"userId"`      //用户Id
	CreatedAt   time.Time `json:"createdAt"`   //创建时间
}
