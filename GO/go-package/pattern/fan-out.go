package pattern

import "fmt"

func FanOut() {
	ch := make(chan int, 100)

	for i := 1; i <= 100; i++ {
		go fanOut(i, ch)
	}

	for i := 0; i < 100000000; i++ {
		ch <- i
	}
	close(ch) // Close the channel when all values are sent

}

func fanOut(workId int, ch <-chan int) {
	var sum int
	for c := range ch {
		sum += c
		fmt.Println(workId, sum)
	}
}
