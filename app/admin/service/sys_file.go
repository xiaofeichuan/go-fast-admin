package service

import (
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/global"
)

type SysFileService struct{}

// Query 文件管理查询
func (s *SysFileService) Query(query dto.SysFileQuery) (list []dto.SysFileVo, total int64, err error) {
	db := global.DB.Model(&model.SysFile{})

	//查询条件

	//总条数
	db.Count(&total)

	//查询数据
	offset := (query.PageNum - 1) * query.PageSize
	err = db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, err
}

// Add 添加文件管理
func (s *SysFileService) Add(addDto dto.SysFileAddDto) error {
	var model = &model.SysFile{
		StorageType: addDto.StorageType,
		BucketName:  addDto.BucketName,
		FileName:    addDto.FileName,
		FileExt:     addDto.FileExt,
		MimeType:    addDto.MimeType,
		FileSize:    addDto.FileSize,
		FileUrl:     addDto.FileUrl,
		IpAddress:   addDto.IpAddress,
		UserId:      addDto.UserId,
	}
	err := global.DB.Create(model).Error
	return err
}

// Detail 获取文件管理详情
func (s *SysFileService) Detail(id int64) (obj dto.SysFileVo, err error) {
	err = global.DB.Model(&model.SysFile{}).Where("id = ?", id).Scan(&obj).Error
	return obj, err
}
