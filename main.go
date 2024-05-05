package main

import (
	"fmt"
)

func main() {
	var lower_limit, upper_limit int

	fmt.Print("Enter the lower limit: ")
	_, err := fmt.Scanln(&lower_limit)
	if err != nil {
		fmt.Println("Invalid lower limit entry: ", err)
		return
	}

	fmt.Print("Enter the upper limit: ")
	_, err = fmt.Scanln(&upper_limit)
	if err != nil {
		fmt.Println("Invalid upper limit entry: ", err)
		return
	}

	//Limit control.
	if lower_limit > upper_limit {
		fmt.Println("The lower limit cannot be greater than the upper limit.")
		return
	}

	//Finding prime number.
	fmt.Println("Prime numbers: ")
	for i := lower_limit; i <= upper_limit; i++ {
		if is_prime(i) {
			fmt.Print(i, ", ")
		}
	}
}

func is_prime(number int) bool {
	//If the number is less than 1, it is not a prime number.
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number % i == 0 {
			return false
		}
	}

	return true
}
