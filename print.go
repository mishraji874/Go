package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")

  fmt.Print(i, "\n",j)

  fmt.Print(i, " ", j)

  var a,b = 10,20

  fmt.Print(a,b)

  var c string = "Hello"
  var d int = 15

  fmt.Printf("c has value: %v and type: %T\n", c, c) //%v is used to print the value of the arguments
  fmt.Printf("d has value: %v and type: %T", d, d) //  %T is used to print the type of the arguments

  
}