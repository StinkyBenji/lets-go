package main

import (
	"fmt"
)

const german = "German"
const spanish = "Spanish"
const englishHello = "Hello, "
const germanHello = "Hallo, "
const spanishHello = "Hola"

// if you have a lot "if" statement, maybe think of using "switch"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// named return value (prefix string): it will be assigned a "zero value"
// calling it by return or return prefix
// publich functions start with a capital letter, private functions start with a lowercase

func greetingPrefix(language string) (prefix string) {
	switch language {
	case german:
		prefix = germanHello
	case spanish:
		prefix = spanishHello
	default:
		prefix = englishHello
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
