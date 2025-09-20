package main

import "fmt"

var lessonNums = map[int]int{
	1: 2,
	2: 4,
	3: 4,
	4: 4,
	5: 2,
	6: 0,
	7: 0,
}

/* При написанні програми використовувати мапи */
/* Написати програму, яка за номером дня тижня (натуральному числу від 1 до 7) видає в якості результату кількість уроків у Вашій групі в цей день */
func main() {
	var day int
	var err error

	fmt.Print("Enter a day number: ")
	_, err = fmt.Scan(&day)
	if err != nil {
		panic(err)
	}

	if val, ok := lessonNums[day]; ok {
		fmt.Printf("Lessons count: %d\n", val)
	} else {
		panic("Invalid day")
	}
}
