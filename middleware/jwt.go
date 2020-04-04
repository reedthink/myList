package middleware

import (
	"myList/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//ReleaseToken 发放token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hhh",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

//claim
//n.索赔；声明；断言
//v.声称；宣称；认领；索取

/*
解码命令 echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 | base64 -d

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
eyJVc2VySWQiOjEwLCJleHAiOjE1ODUzOTUxOTQsImlhdCI6MTU4NDc5MDM5NCwiaXNzIjoiaGhoIiwic3ViIjoidXNlciB0b2tlbiJ9.
zvyhKkYoaghwDnDh3UoPa_d2k3e7bhb_cvidSl41h7U
*/