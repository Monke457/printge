package printge

import "fmt"

type Color string 

const (
	DEFAULT Color = "\033[0;39m"
	BLACK = "\033[0;30m"
	DARKRED = "\033[0;31m"
	DARKGREEN = "\033[0;32m"
	DARKYELLOW = "\033[0;33m"
	DARKBLUE = "\033[0;34m"
	DARKMAGENTA = "\033[0;35m"
	DARKCYAN = "\033[0;36m"
	GRAY = "\033[0;37m"
	DARKGRAY = "\033[1;90m"
	RED = "\033[1;91m"
	GREEN = "\033[1;92m"
	YELLOW = "\033[1;93m"
	BLUE = "\033[1;94m"
	MAGENTA = "\033[1;95m"
	CYAN = "\033[1;96m"
	WHITE = "\033[1;97m"
)

func Print(content string, color Color) {
	fmt.Printf("%s%s%s", color, content, DEFAULT)
}

var list = []Color{RED, YELLOW, GREEN, BLUE, MAGENTA, CYAN, DARKRED, DARKBLUE, DARKCYAN, DARKYELLOW, DARKGREEN, WHITE, BLACK, DARKGRAY, GRAY}
var current = 0

func GetNext() Color {
	next := list[current]
	current++
	if current > len(list) {
		current = 0
	}
	return next
}
