package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 使用bcrypt加密密码
func HashPassword(password string) (string, error) {
	// GenerateFromPassword自动加盐并生成哈希
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash 验证密码是否匹配
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
