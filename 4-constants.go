package main
import (
	"fmt" //importing fmt package
	"math" //importing math package
)

const s string = "constant"

func main(){
	fmt.Println(s)
	//constant
	
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	//6e+11
	fmt.Println(int64(d))
	//600000000000
	fmt.Println(math.Sin(n))
	//0.8250781914

	//A numeric constant has no type until it’s given one, 
	//such as by an explicit conversion.

	//A number can be given a type by using it in a context that requires one, 
	//such as a variable assignment or function call. For example, here math.Sin expects a float64.
}