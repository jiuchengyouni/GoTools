package token

/**
 * token相关的配置信息
 * @author nuonuo
 * @version v1.0
 */

import (
	"strconv"
	"time"
	"github.com/dgrijalva/jwt-go"
)


type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//发布token
func ReleaseToken(uid string) (string, error) {
	uId, _ := strconv.Atoi(uid)
	//设置过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	//创建认证
	claims := &Claims{
		UserId: uint(uId),
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expirationTime.Unix(),
			//发放时间
			IssuedAt: time.Now().Unix(),
			//发放者
			Issuer: "sends",
			//主题
			Subject: "login token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenstring, nil

}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}