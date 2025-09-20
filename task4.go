package main

import (
	"fmt"
	"math/rand"
	"slices"
)

/* Створити програму обробки одновимірного масиву */
/* Видалення - Максимальний елемент */
/* Додавання - К елементів в початок масива */
/* Перестановка - Перевернути масив */
/* Пошук - перший парний */
func main() {
	var err error
	var arr []int
	var initialSize int

	fmt.Print("Enter initial array size: ")
	_, err = fmt.Scan(&initialSize)
	if err != nil {
		panic(err)
	}

	arr = make([]int, initialSize)
	for i := 0; i < initialSize; i++ {
		arr[i] = rand.Intn(128)
	}

	fmt.Printf("Array: %v\n", arr)

	var maxIndex int
	for i := 0; i < initialSize; i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}

	fmt.Printf("Max element: %d\n", arr[maxIndex])

	arr = append(arr[:maxIndex], arr[maxIndex+1:]...)

	fmt.Printf("Array after removing max element: %v\n", arr)

	var addElementsCount int
	fmt.Print("Enter K: ")
	_, err = fmt.Scan(&addElementsCount)
	if err != nil {
		panic(err)
	}

	for i := 0; i < addElementsCount; i++ {
		var num int
		fmt.Printf("Enter number #%d: ", i)
		_, err = fmt.Scan(&num)
		if err != nil {
			panic(err)
		}

		arr = slices.Insert(arr, i, num)
	}

	fmt.Printf("Array after adding %d numbers at the start: %v\n", addElementsCount, arr)

	arrLen := len(arr)
	for i := 0; i < arrLen/2; i++ {
		last := arrLen - i - 1
		arr[i], arr[last] = arr[last], arr[i]
	}

	fmt.Printf("Array after reversing: %v\n", arr)

	firstEvenIndex := -1
	for i := 0; i < arrLen; i++ {
		if arr[i]%2 == 0 {
			firstEvenIndex = i
			break
		}
	}

	if firstEvenIndex >= 0 {
		fmt.Printf("First even element: %d\n", arr[firstEvenIndex])
	} else {
		fmt.Printf("There are no even elements in the array\n")
	}
}
