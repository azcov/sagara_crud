package util

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

//HashPassword will generate hashed password
func HashPassword(password string) (string, error) {
	zap.S().Info("hasing")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	zap.S().Info(bytes, err)
	return string(bytes), err
}

//CheckPasswordHash will
func CheckPasswordHash(hash, password string) error {
	zap.S().Infof("compare hash %s == %s", hash, password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
