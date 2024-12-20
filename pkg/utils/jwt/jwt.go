package jwt

import (
	"auth/pkg/ent"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId  int     `json:"userId"`
	Phone   string  `json:"phone"`
	RoleIds []int64 `json:"roleIds"`
	jwt.StandardClaims
}

const (
	UserLoginKeepAliveTime = time.Hour * 720         // 用户保持登录时间
	UserSignHmacKey        = "LDX_HDZ_USER_SIGN_KEY" // 用户签名key
	UserSessionKey         = "user_session"
)

func GenerateToken(user *ent.User, roleIds []int64) (string, error) {
	claims := Claims{
		UserId:  user.ID,
		Phone:   user.Phone,
		RoleIds: roleIds,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(UserLoginKeepAliveTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(UserSignHmacKey))
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(UserSignHmacKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
