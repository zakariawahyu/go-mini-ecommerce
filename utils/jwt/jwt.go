package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"go-mini-ecommerce/config"
	"time"
)

const (
	TokenExpiredTime = 5 * 60 * 60 // 5 hours
)

func GenerateToken(payload map[string]interface{}) (string, error) {
	cfg, _ := config.LoadConfig()
	tokenClaims := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * TokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenClaims)
	token, err := jwtToken.SignedString([]byte(cfg.Jwt.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
