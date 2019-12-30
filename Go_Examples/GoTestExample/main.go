package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT() (string,err) {
	token := jwt.New()
}

func main(){
	fmt.Println("Hi Go Again")
}