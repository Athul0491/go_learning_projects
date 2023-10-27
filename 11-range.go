//range is used to iterate over elements in a variety of data structures.

package main
import "fmt"

func main() {

	nums := []int{2,3,4}
	sum := 0
	for _,num := range nums {
		sum += num
	}
	fmt.Println("sum:",sum)

	//range on arrays and slices provides both the index and value for each entry. 
	//Above we didn’t need the index, so we ignored it with the blank identifier _. 
	//Sometimes we actually want the indexes though.

	for i,num := range nums {
		if num == 3 {
			fmt.Println("index:",i)
		}
	}
	//output index: 1

	//range on map iterates over key/value pairs.
	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n",k,v)
	}
	//output a -> apple 
	//b -> banana

	//range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:",k)
	}
	//output key: a
	//key: b

	//range on strings iterates over Unicode code points.
	for i,c := range "go" {
		fmt.Println(i,c)
	}
	//output 0 103
	//1 111
	//here the first value is the starting byte index of the rune and the second the rune itself.
}