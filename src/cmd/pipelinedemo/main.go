package main

import (
	"pipeline"
	"fmt"
)

func main() {
	//p := pipeline.ArraySource(3, 2, 6, 7, 4)
	//for {
	//	if num, ok := <- p; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}
	// 和上面方式等价

	p := pipeline.Merge(
			pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
			pipeline.InMemSort(pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))
	for v := range p {
		fmt.Println(v)
	}
}
