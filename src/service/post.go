package service

import (
	"context"
	"github.com/DopamineNone/bubblePro/src/dao/db"
	"github.com/DopamineNone/bubblePro/src/infra/db/mysql"
	"github.com/DopamineNone/bubblePro/src/model"
	idGen "github.com/DopamineNone/bubblePro/src/utils/id_generator"
)

func CreatePost(c context.Context, param *model.Post) error {
	access := db.NewDataAccess(c, mysql.DB)
	param.PostID = idGen.GetID()
	return access.CreatePost(param)
}

func GetPost(c context.Context, pid int64) (post *model.Post, err error) {
	access := db.NewDataAccess(c, mysql.DB)
	return access.GetPostByID(pid)
}
