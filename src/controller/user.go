package controller

import (
	"github.com/DopamineNone/bubblePro/src/service"
	"github.com/DopamineNone/bubblePro/src/utils/echo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SignUpParams struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"eqfield=Password"`
}

type SignInParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshParam struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// AuthResponse reponse of login, sign up, and refresh token
type AuthResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func SignUpHandler(c *gin.Context) {
	// extract param
	param := new(SignUpParams)
	if err := c.ShouldBindJSON(param); err != nil {
		// error log
		zap.L().Error("SignUpHandler - ShouldBindJSON err", zap.Error(err))
		echo.SendErrorResponse(c, err.Error())
		return
	}

	// validate
	if len(param.Username) == 0 || len(param.Password) == 0 || param.ConfirmPassword != param.Password {
		zap.L().Error("SignUpHandler - invalid param")
		echo.SendErrorResponse(c, "password not confirmed")
		return
	}
	// call service
	if accessToken, refreshToken, err := service.SignUpService(c, param.Username, param.Password); err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	} else {
		echo.SendSuccessResponse(c, AuthResponse{
			Access:  accessToken,
			Refresh: refreshToken,
		})
	}
}

func SignInHandler(c *gin.Context) {
	param := new(SignInParams)

	// get param
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("SignInHandler - ShouldBindJSON", zap.Error(err))
		echo.SendErrorResponse(c, err.Error())
		return
	}

	// call service
	if accessToken, refreshToken, err := service.SignInService(c, param.Username, param.Password); err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	} else {
		echo.SendSuccessResponse(c, AuthResponse{
			Access:  accessToken,
			Refresh: refreshToken,
		})
	}
}

func RefreshHandler(c *gin.Context) {
	param := new(RefreshParam)
	if err := c.ShouldBindJSON(param); err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	}
	if accessToken, refreshToken, err := service.RefreshService(c, param.Access, param.Refresh); err != nil {
		echo.SendErrorResponse(c, err.Error())
		return
	} else {
		echo.SendSuccessResponse(c, AuthResponse{
			Access:  accessToken,
			Refresh: refreshToken,
		})
	}
}
