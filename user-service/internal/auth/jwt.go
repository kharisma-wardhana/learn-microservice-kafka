package auth

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kharisma-wardhana/learn-microservice-kafka/user-service/internal/model"
)

type Claims struct {
	UserID string     `json:"user_id"`
	Role   model.Role `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret []byte
var ttlMinutes int

func Init() error {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return errors.New("JWT Secret Not Set")
	}
	jwtSecret = []byte(secret)

	ttl := os.Getenv("JWT_TTL_MINUTES")
	if ttl == "" {
		return errors.New("JWT TTL Not Set")
	} else {
		v, err := strconv.Atoi(ttl)
		if err != nil {
			ttlMinutes = 60
		} else {
			ttlMinutes = v
		}
	}
	return nil
}

func GenerateToken(userID string, role model.Role) (string, error) {
	now := time.Now().UTC()
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ttlMinutes) * time.Minute)),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "user-service",
			Subject:   userID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

type ctxKey string

const ContextClaimsKey ctxKey = "jwt_claims"

func FromContext(ctx context.Context) (*Claims, bool) {
	c, ok := ctx.Value(ContextClaimsKey).(*Claims)
	return c, ok
}
