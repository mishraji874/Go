package main
import (
	"fmt" 
	"math"
)

func main() {
	var num float64 = 24
	fmt.Println("Telusko")
	var result = math.Sqrt(num)

	var intResult = math.Round(result) //we can use Round() to round-up the value
	//use ceil() to round-up the value to the upper level
	//use floor() to round-up the value to the lower level
	
	fmt.Printf("The output is %g",result)
	fmt.Printf("The output is %.2f",intResult)
}