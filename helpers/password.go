package helpers

import "golang.org/x/crypto/bcrypt"

func CheckPassword(Hash string, Password []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(Hash), Password); err != nil {
		return false, err
	}
	return true, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
