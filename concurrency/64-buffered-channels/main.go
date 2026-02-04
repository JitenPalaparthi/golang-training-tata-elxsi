package main

func main() {

	ch1 := make(chan int, 2) // buffered channel

	ch1 <- 100
	ch1 <- 200

	println(<-ch1)
	println(<-ch1)
	ch1 <- 300
	println(<-ch1)
}

// Bufferend channel
// The sender is not block
