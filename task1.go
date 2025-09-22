package main

import "fmt"

func panic_(v any) bool {
	panic(v)
	return true
}

/* Використовуючи логічні вирази вирішити задачу */
/* Написати програму знаходження суми більшого та меншого із трьох чисел */
func main() {
	var err error
	var num1, num2, num3 int

	fmt.Print("Enter a number #1: ")
	_, err = fmt.Scan(&num1)
	_ = err != nil && panic_(err)

	fmt.Print("Enter a number #2: ")
	_, err = fmt.Scan(&num2)
	_ = err != nil && panic_(err)

	fmt.Print("Enter a number #3: ")
	_, err = fmt.Scan(&num3)
	_ = err != nil && panic_(err)

	var min_, max_ int

	min_ = num2 ^ ((num1 ^ num2) & -((num1 - num2) >> 31 & 1))
	min_ = num3 ^ ((min_ ^ num3) & -((min_ - num3) >> 31 & 1))

	max_ = num1 ^ ((num1 ^ num2) & -((num1 - num2) >> 31 & 1))
	max_ = max_ ^ ((max_ ^ num3) & -((max_ - num3) >> 31 & 1))

	// min_ = min(num1, num2, num3)
	// max_ = max(num1, num2, num3)

	/*if num1 > num2 {
		min_ = num2
		max_ = num1
	} else {
		max_ = num2
		min_ = num1
	}
	if min_ > num3 {
		min_ = num3
	}
	if max_ < num3 {
		max_ = num3
	}*/

	fmt.Printf("Sum of min and max is %d", min_+max_)
}
