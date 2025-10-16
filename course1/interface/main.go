package main

type LangBot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (EnglishBot) getGreeting() string {
	return "Hello!"
}

func (SpanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(lb LangBot) {
	println(lb.getGreeting())
}
