// InsertStringSlice 将切片插入到另一个切片的指定位置
package main

import (
	"fmt"
)

func main() {
	src := []string{"a", "b", "c", "d", "e"}
	newstr := []string{"f", "g"}
	fmt.Println(append(newstr, src[:]...))
}
