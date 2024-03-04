package main
import "fmt"

type Student struct {
	rollno int
	name string
	marks int
}
func main() {
	var s1 = Student{101, "Navin", 55}
	fmt.Println(s1)
	fmt.Println(s1.name)

	var s2 = Student{102, "Aditya", 99}
	fmt.Println(s2)
	fmt.Println(s2.marks)

	var s3 = Student{rollno: 103, marks: 65}
	fmt.Println(s3)
	fmt.Println(s3.name)
}