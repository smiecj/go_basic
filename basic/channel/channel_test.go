package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(2 * time.Second)
}

func sendData(ch chan string) {
	ch <- "golang"
}

func getData(ch chan string) {
	fmt.Println(<-ch)
}
