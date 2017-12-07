package main // if 和 for 的风格类似，条件语句不用括号()，执行必须 {}包起来
import "fmt"

func isAdult(age int) string {
	if age < 18 {
		return "It's children"
	} else if age < 60 {
		return "It's Adult"
	} else {
		return "It's old man"
	}
}

func main() {
	fmt.Println("Q: Jack 17 year's old, He is Adult?")
	fmt.Println("A:", isAdult(17))
	fmt.Println("Q: Tony 70 year's old, He is Adult?")
	fmt.Println("A:", isAdult(70))
	fmt.Println("Q: Tom 38 year's old, He is Adult?")
	fmt.Println("A:", isAdult(38))

	// If with a short statement
	// Like for, the if statement can start with a short statement to execute before the condition
	// Variable declared by the statement are only in scope until the end of the if

	if v := 1; v > 0 {
		fmt.Printf("v -> %v", v)
	}
}
