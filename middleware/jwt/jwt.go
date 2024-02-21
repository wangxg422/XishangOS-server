package jwt

import (
	"errors"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpiredOrNotValidYet = errors.New("token is expired or not active yet")
	TokenSignatureInvalid     = errors.New("token signature is invalid")
	TokenMalformed            = errors.New("that's not even a token")
	TokenInvalid              = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.AppConfig.Jwt.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	bf, _ := ParseDuration(global.AppConfig.Jwt.BufferTime)
	ep, _ := ParseDuration(global.AppConfig.Jwt.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		// 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		BufferTime: int64(bf / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"XiShangOS"},             // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天
			Issuer:    global.AppConfig.Jwt.Issuer,               // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := utils.GetSingleFlight().Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, TokenMalformed
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, TokenSignatureInvalid
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, TokenExpiredOrNotValidYet
		default:
			return nil, TokenInvalid
		}
	} else {
		if token.Valid {
			if claims, ok := token.Claims.(*CustomClaims); ok {
				return claims, nil
			}
		}
	}
	return nil, TokenInvalid
}
