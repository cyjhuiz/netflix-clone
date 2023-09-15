package authutil

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func ValidatePassword(inputPassword string, actualPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(inputPassword))
}

func EncryptPassword(inputPassword string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}

func CreateJWT(userID int64) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	// Create the Claims
	claims := &jwt.MapClaims{
		"expiresAt": 1500000,
		"userID":    userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func IsActualUser(token *jwt.Token, userID string) bool {
	claims := token.Claims.(jwt.MapClaims)
	claimsIdStr := fmt.Sprintf("%v", claims["userID"])

	return claimsIdStr == userID
}
