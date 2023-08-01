package main

import "fmt"

func main() {
	// create channel to store the result
	sumSquareResult := make(chan int)
	squareSumResult := make(chan int)

	// start go routine
	go SumSquare(100, sumSquareResult)
	go SquareSum(100, squareSumResult)

	// receive the result from channel
	sumOfSquare := <-sumSquareResult
	squareOfSum := <-squareSumResult

	// calculate the difference
	difference := squareOfSum - sumOfSquare

	fmt.Printf("Sum of squares: %d\n", sumOfSquare)
	fmt.Printf("Square of sum: %d\n", squareOfSum)
	fmt.Printf("Diffenrence: %d\n", difference)
}

func SumSquare(num int, result chan int) {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	result <- sum * sum
}

func SquareSum(num int, result chan int) {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i * i
	}
	result <- sum
}
