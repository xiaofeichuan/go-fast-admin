package service

import (
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/global"
	"time"
)

type SysLoginLogService struct{}

// Query 登录日志查询
func (s *SysLoginLogService) Query(query dto.SysLoginLogQuery) (list []dto.SysLoginLogVo, total int64, err error) {
	db := global.DB.Table("sys_login_log AS log").
		Select("log.*, user.user_name").
		Joins("LEFT JOIN sys_user AS user ON log.user_id = user.id")

	//查询条件
	if query.UserId != 0 {
		db = db.Where("log.`user_id` = ? ", query.UserId)
	}

	if query.StartTime != "" && query.EndTime != "" {
		db = db.Where("log.`login_time` BETWEEN ? and ?", query.StartTime, query.EndTime)
	}

	//总条数
	db.Count(&total)

	//查询数据
	offset := (query.PageNum - 1) * query.PageSize
	err = db.Order("log.id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, err
}

// Add 添加登录日志
func (s *SysLoginLogService) Add(addDto dto.SysLoginLogAddDto) error {
	var model = &model.SysLoginLog{
		UserId:    addDto.UserId,
		LoginTime: time.Now(),
		Ip:        addDto.Ip,
		Location:  addDto.Location,
		Browser:   addDto.Browser,
		OS:        addDto.OS,
	}
	err := global.DB.Create(model).Error
	return err
}

// Detail 获取登录日志详情
func (s *SysLoginLogService) Detail(id int64) (obj dto.SysLoginLogVo, err error) {
	err = global.DB.Model(&model.SysLoginLog{}).Where("id = ?", id).Scan(&obj).Error
	return obj, err
}
