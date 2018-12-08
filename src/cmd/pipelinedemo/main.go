package main

import (
	"pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	//const filename = "small.in"
	//const n = 64
	const filename = "large.in"
	const n = 10000000

	// Create file
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write file
	p := pipeline.RandomSource(n)
	// add buffer
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	// must flush
	writer.Flush()

	// Read file
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count ++
		if count == 100 {
			break
		}
	}

}

func mergeDemo() {
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
