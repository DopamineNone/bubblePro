package controller

import (
	"fmt"
	"github.com/DopamineNone/bubblePro/src/model"
	"github.com/DopamineNone/bubblePro/src/service"
	"github.com/DopamineNone/bubblePro/src/utils/echo"
	"github.com/DopamineNone/bubblePro/src/utils/request"
	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	uid, ok := request.MustUserID(c)
	if !ok {
		return
	}
	param := new(model.PostRequest)
	if err := c.ShouldBindJSON(param); err != nil {
		fmt.Println(err.Error())
		echo.SendErrorResponse(c, err.Error())
		return
	}

	newPost := &model.Post{
		AuthorID:    uid,
		CommunityID: param.CommunityID,
		Title:       param.Title,
		Content:     param.Content,
	}
	if err := service.CreatePost(c, newPost); err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	}
	echo.SendSuccessResponse(c, nil)
}

type GetPostQuery struct {
	Pid int64 `uri:"pid" binding:"numeric"`
}

func GetPostHandler(c *gin.Context) {
	param := new(GetPostQuery)
	if err := c.ShouldBindUri(param); err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	}
	post, err := service.GetPost(c, param.Pid)
	if err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	}
	echo.SendSuccessResponse(c, post)
}
