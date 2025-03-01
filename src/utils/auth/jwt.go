package auth

import (
	"github.com/DopamineNone/bubblePro/src/config"
	appErr "github.com/DopamineNone/bubblePro/src/utils/error"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const issuer = "bluebell"

var (
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
	secret                     []byte
)

type Claim struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GetJwtToken generate jwt tokens
func GetJwtToken(userID int64) (string, string, error) {
	accessToken := Claim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken)
	accessTokenString, err := token.SignedString(secret)
	if err != nil {
		return "", "", err
	}
	refreshToken := Claim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(),
			Issuer:    issuer,
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshToken)
	refreshTokenString, err := token.SignedString(secret)
	return accessTokenString, refreshTokenString, err
}

// ParseJwtToken parse token and return claim
func ParseJwtToken(tokenString string) (*Claim, error) {
	claim := new(Claim)
	token, err := jwt.ParseWithClaims(tokenString, claim, getSecretKey)
	if err != nil {
		return nil, &appErr.Error{
			Code:      appErr.UnableToParseTokenErrorCode,
			Message:   appErr.UnableToParseTokenError,
			ExtraInfo: err.Error(),
		}
	}
	if token.Valid {
		return claim, nil
	}
	return nil, &appErr.Error{
		Code:    appErr.InvalidTokenErrorCode,
		Message: appErr.InvalidTokenError,
	}
}

func RefreshToken(accessToken, refreshToken string) (string, string, error) {
	// check if refresh token exceed its deadline
	_, err := ParseJwtToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	// if access token exceed its deadline, then get a new pair of tokens
	claim := new(Claim)
	_, err = jwt.ParseWithClaims(accessToken, claim, getSecretKey)
	if err == nil {
		return accessToken, refreshToken, nil
	}
	if v, ok := err.(*jwt.ValidationError); ok && v.Errors == jwt.ValidationErrorExpired {
		return GetJwtToken(claim.UserID)
	}
	return accessToken, refreshToken, nil
}

func getSecretKey(token *jwt.Token) (interface{}, error) {
	return secret, nil
}

func init() {
	AccessTokenExpireDuration = time.Duration(config.GetConf().Authentication.AccessTokenExpireDuration) * time.Hour
	RefreshTokenExpireDuration = time.Duration(config.GetConf().Authentication.RefreshTokenExpireDuration*12) * time.Hour
	secret = []byte(config.GetEnv("jwt_secret"))
}
