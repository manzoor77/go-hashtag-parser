package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createhashtag(str string) {
	if len(str) > 1 && len(str) <= 140 {
		//Captilize First character of whole string
		capstr := strings.Title(strings.ToLower(str))
		//Remove space with string words
		remspace := strings.ReplaceAll(capstr, " ", "")
		//Check either input string has # symbol at start or not
		//if yes then return true, otherwise return false
		prefix := strings.HasPrefix(remspace, "#")
		if prefix == true {
			fmt.Println(remspace)

		} else {
			fmt.Println("#" + remspace)
		}
	} else {
		fmt.Println("False")
	}
}
func main() {
	//input Random String from user on console
	//Import bufio for Multi word string input from user
	//Becuase Scanf only take sinle word as input from user

	fmt.Println("Enter a String to create a hashtag:")
	inputReader := bufio.NewReader(os.Stdin)
	str, _ := inputReader.ReadString('\n')
	createhashtag(str)
	// Input:  I am Go Developer
	//Output: #IAmGoDeveloper
}
