package main
import "fmt"

func main() {
	num := 9
	if num == 1 {
		fmt.Println("One")
	} else if num == 2{
		fmt.Println("Two")
	} else {
		fmt.Println("None")
	}

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Print("Two")
	default:
		fmt.Println("None")
	}

}