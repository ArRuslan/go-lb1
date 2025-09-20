package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* Створити програму обробки рядка */
/* Надрукуйте слова, в яких трапляється м'який знак */
func main() {
	fmt.Print("Enter a text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	what := strings.Split(str, " ")
	for _, substr := range what {
		if strings.Contains(substr, "ь") || strings.Contains(substr, "Ь") {
			fmt.Println(substr)
		}
	}
}
