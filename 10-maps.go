package main
import (
	"fmt"
	"maps"
)

package main(){
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	//Output: map: map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))
	//Output: len: 2

	delete(m, "k2")
	fmt.Println("map:", m)
	//Output: map: map[k1:7]

	clear(m)
	fmt.Println("map:", m)
	//Output: map: map[]

	//The optional second return value when getting a value from a map 
	//indicates if the key was present in the map.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	//Output: prs: false

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	//Output: map: map[bar:2 foo:1]

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(m, n2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
	//Output: Maps are equal
}