package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/Murolando/hakaton_final_api/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repo *repository.Repository
}

type tokenClaims struct {
	jwt.StandardClaims
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
func (s AuthService) SignIn(mail *string, password *string) (int64, error) {
	pass := s.generateHashPassword(password)
	id, err := s.repo.GetUserByLoginAndPassword(mail, &pass)
	if err != nil {
		return 0, err
	}

	return id, nil

}
func (s *AuthService) SignUp(user ent.User) (map[string]interface{}, error) {
	line := s.generateHashPassword(user.PasswordHash)
	user.PasswordHash = &line

	id, err := s.repo.SignUp(user)
	if err != nil {
		return nil, err
	}
	token, err := s.GenerateToken(id)
	if err != nil {
		return nil, err
	}
	refresh,err := s.NewRefreshToken(id)
	if err != nil{
		return nil,err
	}
	return map[string]interface{}{"token": token,"refresh":refresh}, nil
}

// Для генерации  хэша пароля
func (s *AuthService) generateHashPassword(password *string) string {
	hash := sha1.New()
	hash.Write([]byte(*password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}

// Для генерации jwt токена
func (s *AuthService) GenerateToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   strconv.FormatInt(id, 10),
		},
	})
	str, err := token.SignedString([]byte(os.Getenv("SIGNINKEY")))
	return str, err
}

// Распарсивает jwt токен, для проверки его валидности
func (s *AuthService) ParseToken(accesstoken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid string method")
		}
		return []byte(os.Getenv("SIGNINKEY")), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		id, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			return 0, errors.New("Troubles with convert string to int")
		}
		return id, nil
	}
	return 0, errors.New("token claims are not of type *tokenClaims")
}

// create refresh token
func (s *AuthService) NewRefreshToken(id int64) (string, error) {
	b := make([]byte, 32)

	st := rand.NewSource(time.Now().Unix())
	r := rand.New(st)

	if _, err := r.Read(b); err != nil {
		return "", err
	}
	refresh := fmt.Sprintf("%x", b)
	t := time.Now().Add(4380 * time.Hour)
	expiredAt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	if err := s.repo.SetSession(id, refresh, expiredAt); err != nil {
		return "", err
	}
	return refresh, nil
}

// get userId by Refresh
func (s *AuthService) GetByRefreshToken(refresh string) (int64, error) {
	id, err := s.repo.GetByRefreshToken(refresh)
	if err != nil {
		return 0, nil
	}
	return id, err

}
