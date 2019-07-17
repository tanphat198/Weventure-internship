package GoLearning

const englishHelloPrefix = "Hello , "
const frenchHelloPrefix = "bubaka "
const spanishHelloPrefix = "Hola "
const vietnamHelloPrefix  = "Chao "
const French = "French"
const Spanish  = "Spanish"
const VietNam  = "VietNam"


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language{
	case French:
		prefix = frenchHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	case VietNam:
		prefix = vietnamHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

