package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义秘钥
var MySecret = []byte("welcome peanut go-jwt world")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	jwt.StandardClaims
	Uid      string `json:"uid"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// 然后我们定义JWT的过期时间，这里以2小时为例：
const TokenExpireDuration = time.Hour * 2

func GenerateToken(uid, username, passport string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "peanut-go-jwt-test",                       // 签发人
			IssuedAt:  time.Now().Unix(),
		},
		uid,
		username,
		passport,
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 更新token
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return GenerateToken(claims.Uid, claims.UserName, claims.Password)
	}
	return "", errors.New("invalid token")
}
