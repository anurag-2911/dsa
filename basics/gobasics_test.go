package basics

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestTypeAssertion(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	fmt.Print()
	var i interface{} = "hallo"
	val, ok := i.(string)
	if ok {
		fmt.Println(val)
	}
	var k interface{} = "Hallo"
	v, ok := k.(int)
	if ok {
		fmt.Println(v)
	}

}
func TestGo01(t *testing.T) {
	concurrent := Concurrent{}
	go concurrent.loopt("Hallo", 10)
	concurrent.loopt("Hello", 10)

}

type Concurrent struct {
}

func (c *Concurrent) hello(s string) string {
	return s
}
func (c *Concurrent) loopt(msg string, counter int) {
	for i := 0; i < counter; i++ {
		fmt.Print(c.hello(msg), i)
		fmt.Print(" ")
		time.Sleep(1 * time.Second)
	}
}

func TestChannel01(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- "hallo"
	}()
	result := <-ch
	fmt.Println(result)
}
func TestBufferedChann(t *testing.T) {
	ch := make(chan string, 2)
	ch <- "Hallo"
	ch <- "Welt"
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("I am done")
}
func TestGoSelect(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "Hallo Welt"

	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello World"

	}()

	select {
	case result := <-ch1:
		{
			fmt.Println(result)
		}
	case result := <-ch2:
		{
			fmt.Println(result)
		}
	}

	fmt.Println("all done")
}

/*
Simple Goroutine Task:

Write a program that starts 10 goroutines.
Each goroutine should generate a random number and print it to the console.
Use channels to synchronize the completion of all goroutines.
*/

func TestSimpleGoRoutineTask(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered ", r)
		}
	}()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go genRandomNumber(&wg)
	}
	wg.Wait()
	fmt.Println("all tasks done")
}
func genRandomNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	randomNum := rand.Intn(100)
	fmt.Println(randomNum)
}
func TestSimpleGoRoutineTaskUsingChannels(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered ", r)
		}
	}()
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go genRanNum(&wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Print(val)
		fmt.Print(" ")
	}
	fmt.Println("all tasks are done")
}
func genRanNum(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- rand.Intn(100)
}

/*
Producer-Consumer Problem:

Implement the classic producer-consumer scenario where you have two goroutines communicating through a channel.
One goroutine (the producer) generates data and sends it over a channel,
while the other goroutine (the consumer) receives data from the channel and processes it.
*/

func TestProdCons(t *testing.T) {
	pc := ProdCons{}
	ch:=make(chan int)
	go pc.produce(ch)
	go pc.consume(ch)
	time.Sleep(time.Second)
	fmt.Println(" all tasks done")
}

type ProdCons struct{}

func (pc *ProdCons) produce(ch chan int) {
	ch <- rand.Intn(100)
	close(ch)
}
func (pc *ProdCons) consume(ch chan int) {
	fmt.Print(<-ch)
	fmt.Print()
}
