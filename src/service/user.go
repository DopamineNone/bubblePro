package service

import (
	"context"
	"errors"
	"github.com/DopamineNone/bubblePro/src/config"
	"github.com/DopamineNone/bubblePro/src/dao/db"
	"github.com/DopamineNone/bubblePro/src/infra/db/mysql"
	"github.com/DopamineNone/bubblePro/src/model"
	"github.com/DopamineNone/bubblePro/src/utils/auth"
	"github.com/DopamineNone/bubblePro/src/utils/crypto"
	appErr "github.com/DopamineNone/bubblePro/src/utils/error"
	idGen "github.com/DopamineNone/bubblePro/src/utils/id_generator"
	"gorm.io/gorm"
)

var (
	security config.Security
)

func init() {
	security = config.GetConf().Security
}

// SignUpService check and regiser a new user
func SignUpService(ctx context.Context, username, password string) (accessToken, refreshToken string, err error) {
	access := db.NewDataAccess(ctx, mysql.DB.Session(&gorm.Session{}))
	// check if user id exist
	if err = access.CheckIfUserExists(username); err != nil {
		return
	}
	// generate id
	newID := idGen.GetID()

	// get security config and crypt
	hash, err := crypto.GenerateHashedString([]byte(password), security.HashLength, security.SaltLength, security.Iterations)
	if err != nil {
		return "", "", &appErr.Error{
			Code:      appErr.UnableToGenerateCipherErrorCode,
			Message:   appErr.UnableToGenerateCipherError,
			ExtraInfo: err.Error(),
		}
	}

	// store user infomation (with password hashed)
	err = access.CreateUser(&model.User{
		ID:       newID,
		Username: username,
		Password: hash,
	})
	if err != nil {
		return
	}
	return auth.GetJwtToken(newID)
}

func SignInService(c context.Context, username, password string) (accessToken, refreshToken string, err error) {
	access := db.NewDataAccess(c, mysql.DB)
	// query user info
	user, err := access.GetUserByName(username)
	if err != nil {
		return "", "", err
	}

	// verify password
	if !crypto.VerifyIfPlainAndHashMatched([]byte(password), user.Password, security.HashLength, security.Iterations) {
		return "", "", errors.New("unmatched password")
	}

	// generate auth token
	return auth.GetJwtToken(user.ID)
}

func RefreshService(c context.Context, accessToken, refreshToken string) (string, string, error) {
	return auth.RefreshToken(accessToken, refreshToken)
}
