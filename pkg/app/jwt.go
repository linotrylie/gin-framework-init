package app

import (
	"encoding/base64"
	"encoding/json"
	"equity/core/model/authmodel/response"
	"equity/global"
	"equity/utils/passwordutil"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Uid       int32  `json:"uid"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GetJwtPayload 获取 JWT 载荷信息
func GetJwtPayload(jwtToken string) (*response.AdminJwtRes, error) {
	jwtStr := strings.Split(jwtToken, ".")
	if len(jwtStr) != 3 {
		msg := fmt.Sprintf("jwtToken不符合三段式格式,jwtToken=%s", jwtToken)
		return nil, errors.New(msg)
	}
	decodeString, err := base64.RawStdEncoding.DecodeString(jwtStr[1])
	if err != nil {
		return nil, err
	}
	var res response.AdminJwtRes
	err = json.Unmarshal(decodeString, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GenerateToken(appKey, appSecret string, uid int32) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    passwordutil.Md5Encode(appKey),
		AppSecret: passwordutil.Md5Encode(appSecret),
		Uid:       uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		Claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return Claims, nil
		}
	}
	return nil, err
}
