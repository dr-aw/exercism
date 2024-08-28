/*
	The Collatz Conjecture or 3x+1 problem can be summarized as follows:
	Take any positive integer n. If n is even, divide n by 2 to get n / 2. If n is odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely. The conjecture states that no matter which number you start with, you will always reach 1 eventually.
	Given a number n, return the number of steps required to reach 1.

	For input n = 12, the return value would be 9.
*/

package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (i int, err error) {
	if n < 1 {
		err = fmt.Errorf("the number might be at least 1, given %d", n)
	} else {
		for n != 1 {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = multiply(n)
			}
			i++
		}
	}
	return i, err
}

func multiply(x int) int { return x*3 + 1 }
