package main
import (
	"fmt" //importing fmt package
	"slices" //importing slices package
)

// slices as similar to vectors in C++

func main(){
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	//OUTPUT uninit: [] true true

	//To create an empty slice with non-zero length, use the builtin make. 
	s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	//OUTPUT emp: [  ] len: 3 cap: 3

	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	//OUTPUT set: [a b c] get: c len: 3

	//append is a bit different than other languages
	s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)
	//OUTPUT apd: [a b c d e f]

	//Slices can also be copy’d.  
	c := make([]string, len(s)) //Here we create an empty slice c of the same length as s
    copy(c, s) //and we copy into c from s.
    fmt.Println("cpy:", c)

	//Slices support a “slice” operator with the syntax slice[low:high].
	l := s[2:5]	//excluding the element at index 5
    fmt.Println("sl1:", l)
	//OUTPUT sl1: [c d e]

	l = s[:5] //excludes the element at index 5
    fmt.Println("sl2:", l)

	l = s[2:] //includes the element at index 2
    fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
	//OUTPUT dcl: [g h i]

	//methods of slices
	t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }
	//OUTPUT t == t2

	//slices can be composed into multi-dimensional data structures.
	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
	//OUTPUT 2d:  [[0] [1 2] [2 3 4]]
}