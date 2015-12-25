/* Largest prime factor
Problem 3
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

Answer:
6857
*/

package main

import (
	"fmt"
	"math"
)

const (
	N int64 = 600851475143
	// N int64 = 13195
)

func main() {
	var i int64 = int64(math.Sqrt(float64(N)))
	for ; i > 1; i-- {
		if N%i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}

func isPrime(n int64) bool {
	var i int64 = 2
	for ; i < n; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}
