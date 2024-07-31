package main

import (
	"fmt"
	"strings"
)

//Checking if a string is a palindrome
func isPalindrome(input string) bool {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
  
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
  
  //Accepting valid input
	for {
		fmt.Println("Enter a string to check if it's a palindrome:")
		_, err := fmt.Scanln(&input)
		if err != nil || strings.TrimSpace(input) == "" || strings.ContainsAny(input, "0123456789`~!@#$%^&*()-_=+[]{}\\|;:'\",.<>?/") {
			fmt.Println("Invalid input. Please enter a non-empty single string with alphabetic characters only.")
			continue
		}
		break
	}
  
  //Printing the result
	if isPalindrome(input) {
		fmt.Printf("'%s' is a palindrome!\n", input)
	} else {
		fmt.Printf("'%s' is not a palindrome.\n", input)
	}
}
