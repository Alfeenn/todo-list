package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword), err
}

func CheckHashPassword(Hashedpassword, password string) bool {

	return bcrypt.CompareHashAndPassword([]byte(Hashedpassword), []byte(password)) == nil
}
