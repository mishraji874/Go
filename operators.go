package main
import ("fmt")

func main() {
  var a = 15 + 25
  fmt.Println(a)

  var x = 10
  x +=5
  fmt.Println(x)

  var z = 5
  var y = 3
  fmt.Println(z>y) // returns 1 (true) because 5 is greater than 3
}