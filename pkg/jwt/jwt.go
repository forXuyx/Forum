package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	UserID               int64  `json:"userID"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("柠檬鱼")

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		userID,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	var cc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, cc, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return cc, nil
	}
	return nil, errors.New("invalid token")
}
