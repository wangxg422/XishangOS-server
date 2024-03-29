package jwt

import "github.com/golang-jwt/jwt/v5"

// CustomClaims Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UserId   int64
	Username string
	NickName string
}
