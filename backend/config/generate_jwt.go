package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtConfig struct {
	SecretKey string
	Lifetime  time.Duration
}

type JWTClaims struct {
	jwt.MapClaims
	UserID string `json:"user_id"`
	Exp    int64  `json:"exp"`
}

func NewJwtConfig(secretKey string, lifetimeHour int) *JwtConfig {
	return &JwtConfig{
		SecretKey: secretKey,
		Lifetime:  time.Duration(lifetimeHour) * time.Hour,
	}
}

func (c *JwtConfig) GenerateJWT(userId string) (string, error) {
	claims := JWTClaims{
		UserID: userId,
		Exp:    time.Now().Add(c.Lifetime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.SecretKey))
}

func (c *JwtConfig) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
