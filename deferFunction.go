package main
import "fmt"

func main() {
	a()

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func a() {
	fmt.Println("a begins")
	defer b() /// defer can be used to that when the first function stops then the second function should start working
	fmt.Println("a ends")
}

func b() {
	fmt.Println("in b")
}