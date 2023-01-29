package main

import (
	"fmt"
	"time"
)

func belangrijk(wait <-chan any, release chan<- any) {
	<-wait
	close(release)
}

func lalka(wait chan<- any) {
	for i := 0; i < 100; i++ {
		go megarzhomba()
	}
	time.Sleep(1 * time.Second)
	close(wait)
}

func megarzhomba() {
	i := uint64(0)
privet:
	i++
	if i < 1000_000_000_000 {
		time.Sleep(0)
		goto privet
	}
}

func main() {
	wait := make(chan any)
	fmt.Println("legen")
	release := make(chan any)
	go belangrijk(wait, release)
	go lalka(wait)
	<-release
	fmt.Println("dary")
}
