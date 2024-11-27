package main

import (
	"fmt"
	"crypto/rand"
	"flag"
	"math/big"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0123456789"
	specials = "!@#$%^&*()_+-={}[]|:;<>,.?/~`"
)

func generatePassword(length int , useUpper, useNumbers, useSpecials bool)(string, error){
	charset := lowercase

	if useUpper {
		charset += uppercase
	}
	if useNumbers{
		charset += numbers
	}
	if useSpecials{
		charset += specials
	}
	if len(charset) == 0{
		return "", fmt.Errorf("no character set selected")
	}

	password := make([]byte, length)
	for i := 0; i< length; i++{
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "" , fmt.Errorf("error generating radom number : %v", err)
		}
		password[i] = charset[index.Int64()]
	}
	return string(password), nil
}

func main(){
	length := flag.Int("length", 12,"Length of the password")
	useUpper := flag.Bool("upper", true, "Include uppercase letters")
	useNumbers := flag.Bool("numbers", true, "Include numbers")
	useSpecials := flag.Bool("specails", true, "Include special character")
	flag.Parse()

	password, err := generatePassword(*length, *useUpper, *useNumbers, *useSpecials)
	if err!= nil{
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Generated Password: ", password )
}