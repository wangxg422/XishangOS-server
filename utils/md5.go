package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

// EncryptPassword 密码加密
func EncryptPassword(password string) (string, string, error) {
	salt := make([]byte, 10)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", err
	}

	//生成密文
	dk := pbkdf2.Key([]byte(password), salt, 10, 32, sha256.New)
	pw := hex.EncodeToString(dk)
	saltStr := hex.EncodeToString(salt)

	return pw, saltStr, nil
}

// EncryptPasswordWithSalt 用指定盐加密密码
func EncryptPasswordWithSalt(password, salt string) (string, error) {
	saltByte, err := hex.DecodeString(salt)
	if err != nil {
		return "", err
	}

	//生成密文
	dk := pbkdf2.Key([]byte(password), saltByte, 10, 32, sha256.New)
	pw := hex.EncodeToString(dk)

	return pw, nil
}
