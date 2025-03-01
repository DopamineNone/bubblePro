package db

import (
	"github.com/DopamineNone/bubblePro/src/model"
)

var (
	postModel model.Post
)

func (access *DataAccess) CreatePost(post *model.Post) error {
	return access.db.WithContext(access.ctx).Model(&postModel).Create(post).Error
}

func (access *DataAccess) GetPostByID(pid int64) (post *model.Post, err error) {
	post = new(model.Post)
	err = access.db.WithContext(access.ctx).Model(&postModel).Where("post_id = ?", pid).First(post).Error
	return
}
