package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "es"
const french = "fr"

func Hello(name string, lang string) string { // public func
	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string)(prefix string){ // pricate func
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return // automatically return prefix
}

func main() {
	fmt.Println(Hello("Go World", ""))
}
