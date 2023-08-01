package main

import "fmt"

func main() {
	// create channel to store the numbers
	numChannel := make(chan int)

	// create go routine to send 100 numbers
	go func() {
		for i := 0; i < 100; i++ {
			numChannel <- i
		}
		close(numChannel)
	}()

	// receive data from the channel
	sum := 0
	oddCount := 0
	evenCount := 0

	for num := range numChannel {
		result := ""

		if num%15 == 0 {
			result = "FizzBuzz"
		} else if num%3 == 0 {
			result = "Fizz"
		} else if num%5 == 0 {
			result = "Buzz"
		} else {
			result = fmt.Sprintf("%d", num)
		}
		fmt.Println(result)

		sum += num

		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	// print all
	fmt.Println("Sum of all numbers:", sum)
	fmt.Println("Total odd numbers:", oddCount)
	fmt.Println("Total even numbers:", evenCount)
}
