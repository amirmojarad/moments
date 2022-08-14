package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generate hashed password with 16 cost.
func HashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	return string(bytes), err
}

// check plainPassword and hashedPassword are equal or not. returns True when passwords are equal
func CheckPasswordHash(plainPassword, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}
