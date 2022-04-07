package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const french = "French"
const spanish = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) { //  named return value and private func
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return // just return
}

func main() {

}

// In our case we've gone from Hello()
//to Hello("name"), to Hello("name", "French") in small,
//easy to understand steps.