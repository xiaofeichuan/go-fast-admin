package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthClaims struct {
	UserId uint64 `json:"userId"`
	jwt.StandardClaims
}

var secretKey = []byte("RcWJhezAYXjzVzZRmaX8B8uy2U3Hgg2p")

// 生成token
func GenerateToken(userId uint64, expireTime time.Time) (string, error) {

	claim := AuthClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// 签名时间
			IssuedAt: time.Now().Unix(),
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 签名颁发者
			Issuer: "gin-fast-admin",
			// 签名主题
			Subject: "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(secretKey)
}

// 解析token
func ParseToken(tokenStr string) (*jwt.Token, error) {
	// 解析 token string 拿到 token jwt.Token
	return jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(tk *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}
