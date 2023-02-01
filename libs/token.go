package libs

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecret = os.Getenv("JWT_KEYS")

type clims struct {
	User_Id string
	Role    string
	jwt.StandardClaims
}

func NewToken(uuid, role string) *clims {
	return &clims{
		User_Id: uuid,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		},
	}
}

func (c *clims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString([]byte(mySecret))
}

func CekToken(token string) (*clims, error) {
	log.Println(token)
	tokens, err := jwt.ParseWithClaims(token, &clims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecret), nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	claims := tokens.Claims.(*clims)
	return claims, nil
}
