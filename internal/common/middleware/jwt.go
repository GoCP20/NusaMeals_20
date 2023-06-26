package middleware

import (
	"errors"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type JWTProvider struct {
	// Signature Secret Key In here
	JWTSecret string
}

func NewJWTProvider(secret string) *JWTProvider {
	return &JWTProvider{
		JWTSecret: secret,
	}
}

type JWTCustomClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (provider *JWTProvider) CreateToken(username string, role string, id uint) (string, error) {
	claims := JWTCustomClaims{
		UserId:   strconv.Itoa(int(id)),
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := provider.JWTSecret

	return token.SignedString([]byte(secretKey))
}

func (provider *JWTProvider) ExtractToken(c echo.Context) (*JWTCustomClaims, error) {
	tokenString := c.Request().Header.Get("Authorization")

	sanitizedTokenBearer := strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(sanitizedTokenBearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token header")
		}

		jwtSecret := provider.JWTSecret

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		customClaim := JWTCustomClaims{}

		claims := token.Claims.(jwt.MapClaims)
		customClaim.UserId = claims["user_id"].(string)
		customClaim.Username = claims["username"].(string)
		customClaim.Role = claims["role"].(string)

		return &customClaim, nil
	}

	return nil, err
}
