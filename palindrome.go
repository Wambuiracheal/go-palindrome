package main

import (
	"fmt"
	"strings"
)

func Palindrome(s string) bool{
	s =  strings.ToLower(s)
	length := len(s)
	for i := 0; i < length/2; i++{
		if s[i] != s[length-1-i]{
			return false
		}
	}
	return true
}

func main(){
	var input string
	fmt.Println("Enter a letter: ")
	fmt.Scanln(&input)

	if Palindrome(input){
		fmt.Println("palindrome")
	}else{
		fmt.Println("is not palindrome")
	}
}