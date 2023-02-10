package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"sync"
)

var (
	secretKey []byte
	once      sync.Once
)

var (
	ErrUnknownAlg   = errors.New("unknown alg")
	ErrInvalidToken = errors.New("invalid token")
)

type Claims = jwt.MapClaims

func SetSecretKey(key string) {
	once.Do(func() {
		secretKey = []byte(key)
	})
}

func Generate(claims Claims, expiredAt int64) (string, error) {
	claims["exp"] = expiredAt

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := signer.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Verify(tokenString string) (Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnknownAlg
		}

		return secretKey, nil
	})

	if err != nil {
		return Claims{}, err
	}

	claims, ok := token.Claims.(Claims)

	if ok && token.Valid {
		return claims, nil
	}

	return Claims{}, ErrInvalidToken
}
