/*
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred

array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred
*/

package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)

  var arr3 = [...]int{1,2,3}
  arr4 := [...]int{4,5,6,7,8}

  fmt.Println(arr3)
  fmt.Println(arr4)

  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)

  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])

  price := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(price)

  arr5 := [5]int{1:10,2:40}

  fmt.Println(arr5)

  arr6 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr7 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr6))
  fmt.Println(len(arr7))
}