package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key used for signing tokens (keep it safe!)
//var jwtSecret = []byte("my_super_secret_key_123")

// CustomClaims defines the structure of JWT claims
type CustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	// Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a signed JWT token
func GenerateJWT(username, email, jwtSecret string) (string, error) {
	claims := CustomClaims{
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // Token expires in 1 hour
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "user-service-app",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	return token.SignedString([]byte(jwtSecret))
}

// ValidateJWT parses and validates a JWT token
func ValidateJWT(tokenString, jwtSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims if token is valid
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
