package main

import "fmt"

const (
	englishPrefix = "Hello, "
	spansihPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
	world         = "world"
)

var languages = []string{"en", "es", "fr"}

func Hello(name string, lang string) string {
	if name == "" {
		name = world
	}
	return getGreetingPrefix(lang) + name
}

func getGreetingPrefix(lang string) (prefix string) {
	switch lang {
	case languages[1]:
		prefix = spansihPrefix
	case languages[2]:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello(world, languages[1]))
}
