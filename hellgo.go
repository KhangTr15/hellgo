package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* var locale */
	var greeting string

	/*var languages = [6]string{"en", "es", "de", "fr", "jp", "kr"}
	locale = languages[4] */

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Language: ")
	language, _ := reader.ReadString('\n')
	fmt.Println(language)

	/*
		fmt.Println("Enter language code:")
		fmt.Scanf("%s", &locale)
	*/

	/*
		if locale == "en"{
			greeting = "Hello"
		} else if locale =="es" {
			greetiing = "Hola"
		} else if locale == "de" {
			greeting = "Guten Tag"
		} else{
			greeting = "YO"
		}
	*/

	switch language {
	case "en\n":
		greeting = "Hello"
	case "es\n":
		greeting = "Hola"
	case "de\n":
		greeting = "Guten Tag"
	case "fr\n":
		greeting = "Bonjour"
	case "jp\n":
		greeting = "Ohayo!"
	case "kr\n":
		greeting = "Annyeong!"
	default:
		greeting = "Yo"
	}
	fmt.Printf(greeting + ", Go!\n")
}
