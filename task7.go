package main

import (
	"fmt"
)

type Person struct {
	name    string
	address string
	age     int
}

/* Створити програму обробки масиву структур */
/* Структура - struct person
{
char*name;
char *adres;
int age;
}; */
/* Критерій для пошуку в масиві структур - Імена починаються на лiтеру 'A' */
/* Завдання для обробки масиву рядків (????) - Додати рядок (??????) із заданим номером (????????????????) */
func main() {
	var people []Person
	var err error

	var peopleCount int
	fmt.Printf("People count: ")
	_, err = fmt.Scanf("%d", &peopleCount)
	if err != nil {
		panic(err)
	}

	people = make([]Person, peopleCount)

	for i := 0; i < peopleCount; i++ {
		var name, address string
		var age int

		fmt.Printf("Person #%d name: ", i)
		_, err = fmt.Scanf("%s", &name)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Person #%d address: ", i)
		_, err = fmt.Scanf("%s", &address)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Person #%d age: ", i)
		_, err = fmt.Scanf("%d", &age)
		if err != nil {
			panic(err)
		}

		people[i] = Person{
			name:    name,
			address: address,
			age:     age,
		}
	}

	found := false
	fmt.Println("People with name starting with \"A\":")
	for idx, person := range people {
		if []rune(person.name)[0] == 'A' {
			fmt.Printf("[%d] Name: %s, Address: %s, Age: %d\n", idx, person.name, person.address, person.age)
			found = true
		}
	}

	if !found {
		fmt.Println("(Nothing found)")
	}

	//
}
