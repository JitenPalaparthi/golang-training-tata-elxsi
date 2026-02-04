package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	ch1 := Generate("Generate-1", time.Millisecond*100, time.Second*1)
	ch2 := Generate("Generate-2", time.Millisecond*150, time.Second*2)
	ch3 := Generate("Generate-3", time.Millisecond*200, time.Second*3)
	ch4 := Generate("Generate-4", time.Millisecond*250, time.Second*4)
	outCh := FanIn(ch1, ch2, ch3, ch4)

	for v := range outCh {
		println(v)
	}

}

func Generate(name string, s time.Duration, d time.Duration) chan string {
	ch := make(chan string)
	timeoutCh := time.After(d)
	i := 1

	go func() {
		for {
			select {
			case <-timeoutCh:
				close(ch)
				runtime.Goexit()
			default:
				time.Sleep(s)
				ch <- fmt.Sprint("Goroutine:", name, ">> Value:", i)
			}

			i++
		}
	}()
	return ch
}

func FanIn(chs ...chan string) chan string {
	outCh := make(chan string, len(chs)) // created it as a buffered channel
	go func() {
		wg := new(sync.WaitGroup)
		for _, ch := range chs { // To iterate each chan from a variadic range
			wg.Add(1)
			go func() {
				for v := range ch { // This knows that the channel is closed bez the range loop iterates only until the channel is not closed
					//	println(v)
					// write some logic to take the data and process it
					//slice := strings.Split(v, ">>")
					outCh <- v
				}
				wg.Done()
			}()
		}
		//go func() {
		wg.Wait() // wait is over means all the gorotuines are done
		close(outCh)
		runtime.Goexit()
		//}()
	}()
	return outCh
}

// FanIn --> Data comes thru many channels combined and processed is calleed as FanIn
