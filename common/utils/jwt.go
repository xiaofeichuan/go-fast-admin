package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-fast-admin/server/app/admin/consts"
	"time"
)

type UserAuthClaims struct {
	UserId   int64           //用户Id
	UserName string          //用户账号
	NickName string          //用户昵称
	UserType consts.UserType //用户类型（0普通账号，1超级管理员）
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

// GetUserInfo 获取当前用户信息
func GetUserInfo(c *gin.Context) *UserAuthClaims {
	claims, exists := c.Get("claims")

	userInfo, ok := claims.(*UserAuthClaims)
	fmt.Println(userInfo)
	fmt.Println(exists)
	fmt.Println(ok)
	return userInfo

}

// GetUserId 获取当前用户Id
func GetUserId(c *gin.Context) int64 {
	userInfo := GetUserInfo(c)
	return userInfo.UserId
}

// IsSuperAdmin 是否超级管理员
func IsSuperAdmin(c *gin.Context) bool {
	userInfo := GetUserInfo(c)
	if userInfo.UserType == consts.UserTypeSuperAdmin {
		return true
	}
	return false
}
