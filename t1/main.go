package main

import "fmt"

func comp(num1 int, num2 int) (bool, int) {
	if num1 > num2 {
		return false, num1 - num2
	} else if num2 > num1 {
		return false, num2 - num1
	}
	return true, 0

}
func main() {
	fmt.Println(comp(1, 2))
	fmt.Println(comp(2, 2))
	fmt.Println(comp(7, 6))
}
