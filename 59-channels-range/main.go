package main

import (
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	sig := make(chan struct{})
	go GenereateSq(ch)
	//go Receiver(ch, sig)
	go receiverR(ch, sig)

	<-sig // just a signal to be triggred
}

func GenereateSq(ch chan<- int) { // sender channel
	for i := 1; i <= 10; i++ {
		ch <- i * i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch) // only a sender can close a channel
}

func Receiver(ch <-chan int, sig chan<- struct{}) {
	for {
		v, ok := <-ch
		if ok {
			println(v, ok)
		} else {
			println("channel is closed", ok, v)
			sig <- struct{}{}
			close(sig)
			runtime.Goexit()
		}
	}
}

func receiverR(ch <-chan int, sig chan<- struct{}) {
	for v := range ch { // range loop iterates until the channel is closed
		println(v)
	}
	sig <- struct{}{}
	close(sig)
}
