/* Problem 1 Multiples of 3 and 5 If we list all the natural numbers below 10
that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples
is 23. Find the sum of all the multiples of 3 or 5 below 1000. Answer: 233168
*/

package main

import (
	"fmt"
)

const (
	CEILING = 1000
)

func main() {
	result := 0
	for x := 1; x <= CEILING; x++ {
		if x%3 == 0 || x%5 == 0 {
			result += x
		}
	}
	fmt.Printf("%d\n", result)
}
