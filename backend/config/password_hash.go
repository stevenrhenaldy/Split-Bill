package config

import "golang.org/x/crypto/bcrypt"

const BcryptCost = 12 // 10–14 is typical; 12 is a good default on modern hardware

func HashPassword(plain string) (string, error) {
	// bcrypt handles salt generation internally
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), BcryptCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(hashFromDB, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashFromDB), []byte(plain))
}
