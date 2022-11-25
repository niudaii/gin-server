package utils

import (
	"fmt"
	"github.com/niudaii/gin-server/global"
	"github.com/golang-jwt/jwt"
	"time"
)

type Jwt struct {
	SigningKey  []byte
	ExpiresTime int64
	BufferTime  int64
	Issuer      string
}

var (
	TokenExpired     = fmt.Errorf("token is expired")
	TokenNotValidYet = fmt.Errorf("token not active yet")
	TokenMalformed   = fmt.Errorf("that's not even a token")
	TokenInvalid     = fmt.Errorf("couldn't handle this token")
)

func NewJwt() *Jwt {
	return &Jwt{
		SigningKey:  []byte(global.Server.Jwt.SigningKey),
		ExpiresTime: global.Server.Jwt.ExpiresTime,
		Issuer:      global.Server.Jwt.Issuer,
	}
}

func (j *Jwt) CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,          // 签名生效时间
			ExpiresAt: time.Now().Unix() + j.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    j.Issuer,                          // 签名的发行者
		},
	}
	return claims
}

func (j *Jwt) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *Jwt) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
