package service

import (
	"context"
	"github.com/DopamineNone/bubblePro/src/dao/db"
	"github.com/DopamineNone/bubblePro/src/infra/db/mysql"
	"github.com/DopamineNone/bubblePro/src/model"
	"github.com/gin-gonic/gin"
)

func ListCommunity(ctx context.Context, pageNum int, pageSize int) (list []model.CommunityOverview, err error) {
	access := db.NewDataAccess(ctx, mysql.DB)
	list, err = access.ListCommuninty(pageNum, pageSize)
	return
}

func GetCommunity(ctx *gin.Context, cid int64) (*model.CommunityDetail, error) {
	access := db.NewDataAccess(ctx, mysql.DB)
	return access.GetCommunityByID(cid)
}
