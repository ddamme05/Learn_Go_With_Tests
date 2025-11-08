package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	french             = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
