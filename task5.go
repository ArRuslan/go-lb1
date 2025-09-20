package main

import (
	"fmt"
	"math/rand"
)

type Matrix2d [][]int

func printMatrix(mat Matrix2d) {
	for _, col := range mat {
		for _, value := range col {
			fmt.Printf("%4d ", value)
		}
		fmt.Print("\n")
	}
}

/* Створити програму обробки двовимірного масиву */
/* Обчислити суму позитивних чисел рядка, у яких добуток елементів непарний */
func main() {
	var matSize int
	var err error
	matMin, matMax := -128, 128
	rand.Seed(42)

	fmt.Print("Matrix size: ")
	_, err = fmt.Scan(&matSize)
	if err != nil {
		panic(err)
	}

	mat := make(Matrix2d, matSize)
	for i := 0; i < matSize; i++ {
		mat[i] = make([]int, matSize)
		for j := 0; j < matSize; j++ {
			mat[i][j] = rand.Intn(matMax-matMin) + matMin
		}
	}

	fmt.Println("Matrix:")
	printMatrix(mat)

	for rowIdx, row := range mat {
		isEven := true
		for _, num := range row {
			if (num % 2) != 0 {
				isEven = !isEven
			}
		}

		if isEven {
			continue
		}

		fmt.Printf("Product of elements of row #%d is odd\n", rowIdx)

		sum := 0

		for _, num := range row {
			if num > 0 {
				sum += num
			}
		}

		fmt.Printf("  Sum of positive elements is: %d\n", sum)
	}
}
