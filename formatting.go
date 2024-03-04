package main

import (
	"fmt"
)

func main() {
	var j = 15.5
	var txt = "Hello World!"

	/*	%v	Prints the value in the default format
		%#v	Prints the value in Go-syntax format
		%T	Prints the type of the value
		%%	Prints the % sign */

	fmt.Printf("%v\n", j)
	fmt.Printf("%#v\n", j)
	fmt.Printf("%v%%\n", j)
	fmt.Printf("%T\n", j)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	var i = 15

	/*%b	Base 2
	%d	Base 10
	%+d	Base 10 and always show sign
	%o	Base 8
	%O	Base 8, with leading 0o
	%x	Base 16, lowercase
	%X	Base 16, uppercase
	%#x	Base 16, with leading 0x
	%4d	Pad with spaces (width 4, right justified)
	%-4d	Pad with spaces (width 4, left justified)
	%04d	Pad with zeroes (width 4 */

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)

	/*%s	Prints the value as plain string
	%q	Prints the value as a double-quoted string
	%8s	Prints the value as plain string (width 8, right justified)
	%-8s	Prints the value as plain string (width 8, left justified)
	%x	Prints the value as hex dump of byte values
	% x	Prints the value as hex dump with spaces*/

	var tx = "Hello"

	fmt.Printf("%s\n", tx)
	fmt.Printf("%q\n", tx)
	fmt.Printf("%8s\n", tx)
	fmt.Printf("%-8s\n", tx)
	fmt.Printf("%x\n", tx)
	fmt.Printf("% x\n", tx)

	//%t	Value of the boolean operator in true or false format (same as using %v)

	var a = true
	var b = false

	fmt.Printf("%t\n", a)
	fmt.Printf("%t\n", b)

	/*%e	Scientific notation with 'e' as exponent
	%f	Decimal point, no exponent
	%.2f	Default width, precision 2
	%6.2f	Width 6, precision 2
	%g	Exponent as needed, only necessary digits*/

	var c = 3.141

	fmt.Printf("%e\n", c)
	fmt.Printf("%f\n", c)
	fmt.Printf("%.2f\n", c)
	fmt.Printf("%6.2f\n", c)
	fmt.Printf("%g\n", c)
}
