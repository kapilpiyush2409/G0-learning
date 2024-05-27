package pattern

import (
	"fmt"
)

func FanIn() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 2; i < 4975; i++ {
			ch1 <- 3 * i
		}
	}()
	go func() {
		defer close(ch2)
		for i := 2; i < 4975; i++ {
			ch2 <- 2 * i
		}
	}()
	go func() {
		defer close(ch3)
		for i := 2; i < 4975; i++ {
			ch3 <- 9 * i
		}
	}()

	for {
		product, ok := fanIn(ch1, ch2, ch3)
		if !ok {
			break
		}
		fmt.Println(product)
	}

}

func fanIn(channels ...<-chan int) (int, bool) {
	var values = make([]int, len(channels))
	product := 1
	for i, ch := range channels {
		val, ok := <-ch
		if !ok {
			return 0, false
		}
		values[i] = val

	}
	for _, val := range values {
		product *= val
	}

	return product, true

}
