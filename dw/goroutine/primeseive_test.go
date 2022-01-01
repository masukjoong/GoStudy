package goroutine

import "fmt"

func ExampleGenerate() {
	src := make(chan int)
	go Generate(src)
	for i := 0; i < 5; i++ {
		prime := <-src
		fmt.Println(prime)
		dst := make(chan int)
		go Filter(src, dst, prime)
		src = dst
	}
	// Output:
	// 2
	// 3
	// 5
	// 7
	// 11
}
