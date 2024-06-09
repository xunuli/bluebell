package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// JWT的过期时间
const TokenExpireDuration = time.Hour * 2

// 服务器的密钥
var mySecret = []byte("bluebell.com")

// MyClaims 自定义声明结构并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 因为这里需要记录一个额外username字段，所以要自定义结构体
type MyClaims struct {
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
	//jwt标准的声明
	jwt.RegisteredClaims
}

// GenToken生成JWT
func GenToken(userid int64, username string) (string, error) {
	//创建一个我们自己的声明的结构
	cliams := MyClaims{
		username,
		userid,
		jwt.RegisteredClaims{
			Issuer:    "bluebell",                                              //签发人
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //过期时间
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	//使用指定的secret签名并获得完整的编码后的字符串token
	ss, err := token.SignedString(mySecret)
	return ss, err
}

// ParseToken解析JWT
func ParseToken(tokenstring string) (*MyClaims, error) {
	//解析token
	mc := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenstring, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { //校验token
		return mc, nil
	}
	return nil, nil
}
