package controller

import (
	"fmt"
	"github.com/DopamineNone/bubblePro/src/service"
	"github.com/DopamineNone/bubblePro/src/utils/echo"
	"github.com/gin-gonic/gin"
)

const (
	PageSize10 = 10
	PageSize20 = 20
)

type Pagination struct {
	PageNum  int `form:"page"`
	PageSize int `form:"page_size"`
}

func ListCommunityHandler(c *gin.Context) {
	param := new(Pagination)
	if err := c.ShouldBindQuery(param); err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	}
	// check if page size enum valid
	if param.PageSize != PageSize10 && param.PageSize != PageSize20 {
		echo.SendErrorResponse(c, "invalid page size")
		return
	}

	list, err := service.ListCommunity(c, param.PageNum, param.PageSize)
	if err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	}
	echo.SendSuccessResponse(c, list)
}

type GetCommunityQuery struct {
	Cid int64 `uri:"cid" binding:"numeric"`
}

func GetCommunityHandler(c *gin.Context) {
	param := new(GetCommunityQuery)
	if err := c.ShouldBindUri(param); err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	}
	fmt.Println(param)

	community, err := service.GetCommunity(c, param.Cid)
	if err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	}
	echo.SendSuccessResponse(c, community)
}
