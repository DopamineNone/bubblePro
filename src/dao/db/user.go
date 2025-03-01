package db

import (
	"github.com/DopamineNone/bubblePro/src/model"
	appErr "github.com/DopamineNone/bubblePro/src/utils/error"
)

var userModel = model.User{}

func (access *DataAccess) CheckIfUserExists(username string) error {
	var count int64
	if err := access.db.WithContext(access.ctx).Model(&userModel).
		Where("username = ?", username).Count(&count).Error; err != nil {

		return &appErr.Error{
			Code:      appErr.UnableToQueryUserErrorCode,
			Message:   appErr.UnableToQueryUserError,
			ExtraInfo: err.Error(),
		}
	}
	if count > 0 {
		return &appErr.Error{
			Code:    appErr.UserAlreadyExistedErrorCode,
			Message: appErr.UserAlreadyExistedError,
		}
	}
	return nil
}

func (access *DataAccess) CreateUser(param *model.User) error {
	if err := access.db.WithContext(access.ctx).Create(param).Error; err != nil {
		return &appErr.Error{
			Code:      appErr.UnableToCreateUserErrorCode,
			Message:   appErr.UnableToCreateUserError,
			ExtraInfo: err.Error(),
		}
	}
	return nil
}

func (access *DataAccess) GetUserByName(username string) (user *model.User, err error) {
	user = new(model.User)
	err = access.db.WithContext(access.ctx).Where(&model.User{Username: username}).First(user).Error
	if err != nil {
		return nil, &appErr.Error{
			Code:      appErr.UnableToQueryUserErrorCode,
			Message:   appErr.UnableToQueryUserError,
			ExtraInfo: err.Error(),
		}
	}
	return
}
