package main

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"
	italian = "Italian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	italianHelloPrefix = "Buongiorno, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPreefix(language) + name
}

func greetingPreefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
