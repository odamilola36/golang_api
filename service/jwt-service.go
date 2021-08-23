package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)


type Jwtservice interface{
	GenerateToken(UserID string) string
	ValidateToken(UserID string) (*jwt.Token, error)	
}

type jwtCustomClaim struct{
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer string
}

func NewJWTService() Jwtservice {
	return &jwtService{
		issuer: "Lomari",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		fmt.Printf("secret key is empty, defaulting to %v", "lomari")
	}
	return secretKey
}

func (s jwtService) GenerateToken(UserId string) string {
	
	claims := &jwtCustomClaim{
		UserId,
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			Issuer: s.issuer,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		panic(err)
	}

	return t
}

func (j *jwtService) ValidateToken (token string) (*jwt.Token, error) {
	return jwt.Parse(token, 
		func (s_ *jwt.Token) (interface{}, error ) {
			if _, ok := s_.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, fmt.Errorf("unexpected signing method %v", s_.Header["alg"])
			}
			return []byte(j.secretKey), nil
			// func(*Token) (interface{}, error)
		},
	)
}
