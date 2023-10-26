package main
import "fmt"

func main(){
	//By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("emp:",a)
	//emp: [0 0 0 0 0]

	//inserting an element at a specific index
	a[4]=100
	fmt.Println("set:",a)
	//OUTPUT set: [0 0 0 0 100]

	//getting an element at a specific index
	fmt.Println("get:",a[4])
	//OUTPUT get: 100

	//the builtin len returns the length of an array
	fmt.Println("len:",len(a))
	//OUTPUT len: 5

	//use this syntax to declare and initialize an array in one line
	b:=[5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)
	//OUTPUT dcl: [1 2 3 4 5]

	//array types are one-dimensional, but you can compose types to build multi-dimensional data structures
	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d:",twoD)
	//OUTPUT 2d: [[0 1 2] [1 2 3]]
}