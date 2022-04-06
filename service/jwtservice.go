package service

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(id uint64, username string, iat int64, exp int64) string
	ValidateToken(token string) (*jwt.Token, error)
	GenerateRefreshToken(id uint64, username string, iat int64, exp int64) string
	ValidateRefreshToken(token string) (*jwt.Token, error)
	RenewToken(id uint64, username string, iat int64, exp int64) string
}

type jwtCustomClaim struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type jwtService struct {
	secretAccessKey  string
	secretRefreshKey string
	issuer           string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:           "deliportal",
		secretAccessKey:  getSecretKey(),
		secretRefreshKey: getRefreshTokenSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	if secretKey == "" {
		secretKey = "deliportal"
	}
	return secretKey
}

func getRefreshTokenSecretKey() string {
	secretKey := os.Getenv("JWT_REFRESH_TOKEN_SECRET")
	if secretKey == "" {
		secretKey = "deliportal"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(id uint64, username string, iat int64, exp int64) string {

	claims := &jwtCustomClaim{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    j.issuer,
			IssuedAt:  iat,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretAccessKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretAccessKey), nil
	})
}

func (j *jwtService) GenerateRefreshToken(id uint64, username string, iat int64, exp int64) string {
	claims := &jwtCustomClaim{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    j.issuer,
			IssuedAt:  iat,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretRefreshKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateRefreshToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretRefreshKey), nil
	})
}

func (j *jwtService) RenewToken(id uint64, username string, iat int64, exp int64) string {
	claims := &jwtCustomClaim{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    j.issuer,
			IssuedAt:  iat,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretAccessKey))
	if err != nil {
		panic(err)
	}
	return t
}
