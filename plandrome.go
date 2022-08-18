package main

import "fmt"

func main() {

	result := isPalindrome(121)
	fmt.Println(result)
	//num := 121

}
func isPalindrome(num int) bool {
	a := false
	var temp int = 0
	num2 := num

	for num > 0 {

		temp = (temp * 10) + num%10

		num = num / 10

	}
	if temp == num2 {
		fmt.Println("number is plandrome")
		return true
	} else {

		fmt.Println("number is not plandrome")
		return false

	}

	return a

}
