// Author: magician
// File:   main.go
// Date:   2021/2/26
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string
	jwt.StandardClaims
}

func main()  {
	mySigningKey := []byte("qimiaoshuai")
	// StandardClaims struct
	//c := MyClaims{
	//	Username: "qimiao",
	//	StandardClaims: jwt.StandardClaims{
	//		NotBefore: time.Now().Unix() - 60, // 生效时间
	//		ExpiresAt: time.Now().Unix() + 5, // 失效时间
	//		Issuer:    "qimiao", // 签发人
	//	},
	//}
	//t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// MapClaims map
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Unix() + 5,
		"iss": "qimiao",
		"nbf": time.Now().Unix() - 5,
		"username": "my",
	})
	// sign
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(s)
	//time.Sleep(6 * time.Second)

	// parse
	//token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
	//	return mySigningKey, nil
	//})
	token, err := jwt.ParseWithClaims(s, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	//fmt.Println(token.Claims.(*MyClaims).Username)
	// TODO map指针取值
	fmt.Printf("%T", token.Claims.(*jwt.MapClaims))
	//fmt.Println(&token.Claims["username"])
}
