package main

import "fmt"

// func canIDrink(age int) bool {
// 	if age < 18 {
// 		return false
// 	}
// 	return true
// }

// func canIDrink(age int) bool {
// 	// expression variable
// 	if koreanAge := age + 2; koreanAge < 18 {
// 		return false
// 	}
// 	return true
// }

// func canIDrink(age int) bool {
// 	switch age {    // koreanAge := age + 2; koreanAge {} -> if와 같이 사용가능
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }

func canIDrink(age int) bool {
	switch {
	case age <= 10:
		return false
	case age == 18:
		return true
	case age >= 50:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}
