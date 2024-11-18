package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	JWTService interface {
		GenerateToken(id string) string
		ValidateToken(token string) (*jwt.Token, error)
		GetEntityIdByToken(token string) (string, error)
	}
)

type jwtCustomClaim struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer string
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "Template"
	}
	return secretKey
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer: "Template",
	}
}

func (j *jwtService) GenerateToken(id string) string {
	claims := jwtCustomClaim{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute*60)),
			Issuer: j.issuer,
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return signedToken
}

func (j *jwtService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, j.parseToken)
}

func (j *jwtService) GetEntityIdByToken(token string) (string, error) {
	v_token, err := j.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims := v_token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["id"])
	return id, nil
}
