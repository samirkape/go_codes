//GCD using tuple assignment
package main

//GCD using tuple assignment

import "fmt"

func main() {
	gcd := find_gcd(5, 10)
	fmt.Println(gcd)
}

func find_gcd(num1, num2 int) int {
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	return num1
}
