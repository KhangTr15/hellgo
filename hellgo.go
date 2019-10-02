package main

import "fmt"

func main() {
	var locale, greeting string
	var languages = [6]string{"en", "es", "de", "fr", "jp", "kr"}
	locale = languages[5]

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

	switch locale {
	case "en":
		greeting = "Hello"
	case "es":
		greeting = "Hola"
	case "de":
		greeting = "Guten Tag"
	case "fr":
		greeting = "Bonjour"
	case "jp":
		greeting = "Ohayo!"
	case "kr":
		greeting = "Annyeong!"
	default:
		greeting = "Yo"
	}
	fmt.Printf(greeting + ", Go!\n")
}
