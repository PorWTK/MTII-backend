package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenService interface {
	GenerateToken(userId int) string
	ValidateToken(token string) (*jwt.Token, error)
	InvalidateToken(token string) error
	GetUserIdByToken(token string) (int, error)
}

type CustomClaim struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

type tokenService struct {
	secretKey  string
	issuer     string
	invalidate map[string]time.Time
	mutex      sync.Mutex
}

func NewTokenService() TokenService {
	return &tokenService{
		secretKey:  getSecretKey(),
		issuer:     "Template",
		invalidate: make(map[string]time.Time),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "Template"
	}
	return secretKey
}

func (ts *tokenService) GenerateToken(userId int) string {
	claims := CustomClaim{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			Issuer:    ts.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(ts.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (ts *tokenService) ValidateToken(token string) (*jwt.Token, error) {
	ts.mutex.Lock()
	if _, ok := ts.invalidate[token]; ok {
		ts.mutex.Unlock()
		return nil, fmt.Errorf("token has been invalidated")
	}
	ts.mutex.Unlock()
	return jwt.Parse(token, ts.parseToken)
}

func (ts *tokenService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(ts.secretKey), nil
}

func (ts *tokenService) InvalidateToken(token string) error {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	if _, ok := ts.invalidate[token]; ok {
		return fmt.Errorf("token has already been invalidated")
	}
	ts.invalidate[token] = time.Now()
	return nil
}

func (ts *tokenService) GetUserIdByToken(token string) (int, error) {
	t_Token, err := ts.ValidateToken(token)
	if err != nil {
		return 0, err
	}
	claims := t_Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	userId, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("invalid user id format: %v", err)
	}
	return userId, nil
}
