package auth

import (
	"fmt"
	"github.com/evilbebra/auth/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type Auth interface {
	GenerateJWTToken(user *types.User) (string, error)
	ValidateToken(tokenStr string) (jwt.MapClaims, error)
}

type AuthService struct {
	signingKey     string
	expireDuration time.Duration
}

func NewAuthService(signingKey string, expireDuration int) Auth {
	return &AuthService{
		signingKey:     signingKey,
		expireDuration: time.Duration(expireDuration),
	}
}

func (s *AuthService) GenerateJWTToken(user *types.User) (string, error) {
	claims := jwt.MapClaims{
		"id":        user.ID,
		"email":     user.Email,
		"validTill": time.Now().Add(time.Second * s.expireDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(s.signingKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (s *AuthService) ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unauthorized")
		}

		secret := s.signingKey
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unauthorized: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}

	validTill, ok := claims["validTill"].(float64)
	if !ok {
		return nil, fmt.Errorf("unauthorized: invalid validTill time")
	}

	currentTime := time.Now()
	if currentTime.After(time.Unix(int64(validTill), 0)) {
		log.Println("Token time out")
		return nil, fmt.Errorf("unauthorized: token time out")
	}

	return claims, nil
}
