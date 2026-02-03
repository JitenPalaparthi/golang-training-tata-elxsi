package main

import (
	"runtime"
	"time"
)

func main() {

	ch1 := GenerateEven(time.Millisecond*100, time.Second*1)
	ch2 := GenerateOdd(time.Millisecond*120, time.Second*1)
	done := Receive(ch1, ch2)
	<-done

}

func GenerateEven(s time.Duration, d time.Duration) <-chan int {
	ch := make(chan int)
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
				if i%2 == 0 {
					ch <- i
				}
			}

			i++
		}
	}()
	return ch
}

func GenerateOdd(s time.Duration, d time.Duration) <-chan int {
	ch := make(chan int)
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
				if i%2 != 0 {
					ch <- i
				}
			}

			i++
		}
	}()
	return ch
}

func Receive(ch1, ch2 <-chan int) <-chan struct{} {
	done1, done2 := false, false
	done := make(chan struct{})
	go func() {
		for {
			if done1 && done2 {
				println("finished receiving values")
				done <- struct{}{}
				close(done)
				runtime.Goexit()
			}
			select {
			case v, ok := <-ch1:
				if ok {
					println(v)
				} else {
					//println(v, "ch1 closed stopped receiving")
					done1 = true
				}
			case v, ok := <-ch2:
				if ok {
					println(v)
				} else {
					//println(v, "ch2 closed stopped receiving")
					done2 = true
				}
				// default:
				// 	if done1 && done2 {
				// 		println("finished receiving values")
				// 		done <- struct{}{}
				// 		close(done)
				// 		runtime.Goexit()
				// 	}
			}
		}
		//runtime.Goexit()
	}()
	return done
}
