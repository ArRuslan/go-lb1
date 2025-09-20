package main

import (
	"fmt"
	"math"
)

/* Скласти програму обчислення значення функцій при різних значеннях аргументів, заданих інтервалом зміни і величиною кроку, результат представити у вигляді таблиці */
// y = ax + bcosx for x < 0.5
// y = bx^2 + csin2x for x in [0.5, 1)
// x [0, 2]
// h = 0.1
// a = 0.75
// b = 1.19
// c = -2.5
func main() {
	const a = 0.75
	const b = 1.19
	const c = -2.5

	fmt.Println("+---------+---------+")
	fmt.Println("|    X    |    Y    |")
	fmt.Println("+---------+---------+")

	for x := 0.0; x <= 2; x += 0.1 {
		var y float64

		if x < 0.5 {
			y = a*x + b*math.Cos(x)
		} else if x < 1 {
			y = b*math.Pow(x, 2) + c*math.Sin(2*x)
		}

		fmt.Printf("| %6.3f  | %6.3f  |\n", x, y)
	}

	fmt.Println("+---------+---------+")
}
