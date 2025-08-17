package utils

import "golang.org/x/crypto/bcrypt"

// EncryptPassword 对密码进行加密处理
func EncryptPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bcryptPassword), nil
}

// CheckPassword 对密码进行对比
func CheckPassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
