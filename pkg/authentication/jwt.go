package authentication

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"iago-effting/api-example/configs"
)

type JWTService interface {
	GenerateToken(email string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Identity string `json:"identity"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issuer:    configs.Env.Authentication.Issuer,
	}
}

func getSecretKey() string {
	secret := configs.Env.Authentication.Secret
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(identity string) string {
	timeToExpires := time.Duration(configs.Env.Authentication.Expires) * time.Hour

	claims := &authCustomClaims{
		identity,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeToExpires).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {

		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
