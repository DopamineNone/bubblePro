package db

import "github.com/DopamineNone/bubblePro/src/model"

var commnutiyModel = model.Community{}

func (access *DataAccess) ListCommuninty(pageNum, pageSize int) (list []model.CommunityOverview, err error) {
	err = access.db.WithContext(access.ctx).Model(&commnutiyModel).Find(&list).Error
	return
}

func (access *DataAccess) GetCommunityByID(id int64) (detail *model.CommunityDetail, err error) {
	detail = new(model.CommunityDetail)
	err = access.db.WithContext(access.ctx).Model(&commnutiyModel).First(detail).Error
	return
}
