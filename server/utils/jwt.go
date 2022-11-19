package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

type UserAuthClaims struct {
	UserId   uint64 //用户Id
	Account  string //用户账号
	UserName string //用户名称
	UserType int    //用户类型（0普通账号，1超级管理员）
	jwt.StandardClaims
}

var secretKey = []byte("RcWJhezAYXjzVzZRmaX8B8uy2U3Hgg2p")

// GenerateToken 生成token
func GenerateToken(claim UserAuthClaims, expireTime time.Time) (string, error) {

	claim.StandardClaims = jwt.StandardClaims{
		// 签名时间
		IssuedAt: time.Now().Unix(),
		// 过期时间
		ExpiresAt: expireTime.Unix(),
		// 签名颁发者
		Issuer: "go-fast-admin",
		// 签名主题
		Subject: "admin",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(secretKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*jwt.Token, error) {
	// 解析 token string 拿到 token jwt.Token
	var claim UserAuthClaims
	return jwt.ParseWithClaims(tokenStr, &claim, func(tk *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(c *gin.Context) *UserAuthClaims {
	claims, exists := c.Get("claims")

	userInfo, ok := claims.(*UserAuthClaims)
	fmt.Println(userInfo)
	fmt.Println(exists)
	fmt.Println(ok)
	return userInfo

}

// GetCurrentUserId 获取当前用户Id
func GetCurrentUserId(c *gin.Context) uint64 {
	currentUser := GetCurrentUser(c)
	return currentUser.UserId
}
