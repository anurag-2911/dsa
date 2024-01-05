package gobasics

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestXxxY(t *testing.T) {
	fmt.Println()
}

func TestSlc(t *testing.T) {
	arr := make([]int, 0, 10)
	arr09 := [3]int{0, 1, 2}
	for i := 0; i < 3; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr, arr09)
	changeslc(arr, arr09)
	fmt.Println(arr, arr09)

	list := make(map[int]string)
	list[1] = "one"
	list[20] = "two"
	for i, val := range list {
		fmt.Println(i, val)
	}

	ch := make(chan int)
	go func(ch chan int) {
		ch <- 101
		close(ch)
	}(ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("all done")
}
func changeslc(arr2 []int, arr3 [3]int) {
	arr2[2] = 901
	arr3[2] = 109
}

func TestIXFace(t *testing.T) {
	cl := &ConsoleLogger{}
	logme(cl, "hallo")
}

type logger interface {
	logmessage(msg string)
}

type ConsoleLogger struct {
}

func (cl *ConsoleLogger) logmessage(msg string) {
	fmt.Println(msg)
}
func logme(l logger, msg string) {
	l.logmessage(msg)
}

func TestPR(t *testing.T) {
	pan()
	fmt.Println("all done")

}

func pan() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from a panic ", r)
		}
	}()
	causePanic()
	fmt.Println("after a panic")
}
func causePanic() {
	panic("going to panic")
}

func TestBacicGR(t *testing.T) {
	go sayhallo()
	time.Sleep(1 * time.Second)
	fmt.Println("main routine")
}
func sayhallo() {
	fmt.Println("hallo from go routine")
}

func TestBasicCHN(t *testing.T) {
	ch := make(chan string)
	go sayHola(ch)
	fmt.Println(<-ch)
	fmt.Println("all done")
}
func sayHola(ch chan string) {
	ch <- "hola"
	close(ch)
}
func TestBfChan(t *testing.T) {
	chanlen := 2
	ch := make(chan string, chanlen)
	go multiplemsg(ch, chanlen)
	for val := range ch {
		fmt.Print(val)
		fmt.Print()
	}
	fmt.Println("all done")

}
func multiplemsg(ch chan string, chanlen int) {
	for i := 0; i < chanlen; i++ {
		ch <- fmt.Sprintf("%d", i)
	}
	close(ch)
}

func TestChanSync(t *testing.T) {
	ch := make(chan bool, 1)
	go signal(ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("all done")

}
func signal(ch chan bool) {
	time.Sleep(time.Second * 2)
	ch <- true

	time.Sleep(10 * time.Second)
	fmt.Println("signal routine done")
	close(ch)

}

func TestRangeOnChannel(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for r := range ch {
		fmt.Println(r)
	}
	fmt.Println("all done")
}

func TestSelectMultiChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- "Hallo"
		close(ch1)
	}()
	go func() {
		time.Sleep(time.Second * 5)
		ch2 <- "Hello"
		close(ch2)
	}()

	for ch1open, ch2open := true, true; ch1open || ch2open; {
		select {
		case v1, ok := <-ch1:
			{
				if !ok {
					ch1open = false
					continue
				}
				fmt.Println("received", v1)
			}
		case v2, ok := <-ch2:
			{
				if !ok {
					ch2open = false
					continue
				}
				fmt.Println("received", v2)
			}
		}
	}
	fmt.Println("all done")
}

func TestS2(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "hallo"
		close(ch1)
	}()
	go func() {
		ch2 <- "hello"
		close(ch2)
	}()

	for ch1open, ch2open := true, true; ch1open == true || ch2open == true; {

		select {

		case m1, ok := <-ch1:
			{
				if !ok {
					ch1open = false
				}
				fmt.Println(m1)
			}
		case m2, ok := <-ch2:
			{
				if !ok {
					ch2open = false
				}
				fmt.Println(m2)
			}
		}
	}
	fmt.Println("all done")
}

func TestPassingData(t *testing.T) {
	ch := make(chan int)
	go xsquare(ch)
	ch <- 9
	fmt.Println("result ", <-ch)
}
func xsquare(ch chan int) {
	number := <-ch
	ch <- number * number
	close(ch)
}
func TestProdCons(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go xprod(ch, &wg, i+1)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	readValues(ch)
	fmt.Println("all done")
}

func readValues(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}
func xprod(ch chan<- int, wg *sync.WaitGroup, id int) {
	for i := 0; i < 5; i++ {
		ch <- i * id
	}
	wg.Done()
}

func TestSTimeOut(t *testing.T) {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		ch <- 1001
	}()

	select {
	case val := <-ch:
		{
			fmt.Println("value is ", val)
		}
	case <-time.After(2 * time.Second):
		{
			fmt.Println("time out")
		}
	}

	fmt.Println("all done")
}
