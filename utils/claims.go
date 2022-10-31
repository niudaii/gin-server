package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	Username    string
	AuthorityId string
}

// GetClaims 从 gin 的 context 中获取并解析 jwt
func GetClaims(c *gin.Context) (claims *CustomClaims, err error) {
	token := c.Request.Header.Get("Access-Token")
	j := NewJwt()
	claims, err = j.ParseToken(token)
	return claims, err
}

// GetUserUuid 从 gin 的 context 中获取从 jwt 解析出来的 uuid
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId 从 gin 的 context 中获取从 jwt 解析出来的 AuthorityId
func GetUserAuthorityId(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUsername 从 gin 的 context 中获取从 jwt 解析出来的 username
func GetUsername(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.Username
	}
}
