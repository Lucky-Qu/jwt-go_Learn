package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type MyClaims struct {
	Msg string `json:"msg"`
	jwt.StandardClaims
}

func main() {
	msg := "HelloJwt"
	key := []byte("SomethingSecret")
	Token := encryption(msg, key)
	fmt.Println(Token)
	for i := range 30 {
		fmt.Printf("现在是第%d次解析\n", i+1)
		decryption(Token, key)
		time.Sleep(1 * time.Second)
	}
}

func encryption(msg string, key []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		Msg: msg,
		StandardClaims: jwt.StandardClaims{
			Audience:  "User",
			ExpiresAt: time.Now().Add(time.Second * 20).Unix(),
			Id:        "",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "LuckyQu",
			NotBefore: time.Now().Add(time.Second * 10).Unix(),
			Subject:   "",
		},
	})
	encryptedToken, err := token.SignedString(key)
	if err != nil {
		fmt.Println("生成Token失败")
	}
	return encryptedToken
}
func decryption(token string, key []byte) {
	decryptedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("未经验证的算法")
		}
		return key, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	if claims, ok := decryptedToken.Claims.(jwt.MapClaims); ok && decryptedToken.Valid {
		fmt.Println(claims["msg"])
	}
}
