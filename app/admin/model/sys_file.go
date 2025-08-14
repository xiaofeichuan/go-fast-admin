package model

import (
	"time"
)

type SysFile struct {
	Id          int64     `json:"id" gorm:"primaryKey;comment:文件Id"`
	StorageType string    `json:"storageType" gorm:"comment:存储类型"`
	BucketName  string    `json:"bucketName" gorm:"comment:存储桶"`
	FileName    string    `json:"fileName" gorm:"comment:文件名（如：1.png）"`
	FileExt     string    `json:"fileExt" gorm:"comment:文件扩展名（如： png）"`
	MimeType    string    `json:"mimeType" gorm:"comment:文件 MIME 类型（防止伪装文件上传）"`
	FileSize    int64     `json:"fileSize" gorm:"comment:文件大小（单位：字节）"`
	FileUrl     string    `json:"fileUrl" gorm:"comment:文件地址"`
	IpAddress   string    `json:"ipAddress" gorm:"comment:上传IP"`
	UserId      int64     `json:"userId" gorm:"comment:用户Id"`
	CreatedAt   time.Time `json:"createdAt" gorm:"comment:创建时间"`
}

func (SysFile) TableName() string {
	return "sys_file"
}
