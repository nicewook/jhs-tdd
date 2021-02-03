package hello

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string) string {
	if len(name) == 0 {
		name = "world"
	}
	return englishHelloPrefix + name
}

func HelloInternational(name string, language string) string {
	if len(name) == 0 {
		name = "world"
	}
	if language == "engilsh" {
		return englishHelloPrefix + name
	} else if language == "spanish" {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}
