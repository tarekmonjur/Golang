package main

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		println("Hex code for color", key, "is", value)
	}
}
