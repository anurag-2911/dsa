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
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		pc.produce(ch, 5)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		pc.consume(ch)
	}()
	wg.Wait()
	fmt.Println(" all tasks done")
}

type ProdCons struct{}

func (pc *ProdCons) produce(ch chan int, count int) {
	rand.Seed(50)
	for i := 0; i < count; i++ {
		val := rand.Intn(100)
		fmt.Print(val)
		fmt.Print(" ")
		ch <- val

	}
	close(ch)
}
func (pc *ProdCons) consume(ch chan int) {
	result := 0
	for val := range ch {
		result = result + val
	}
	fmt.Println("result is ", result)
}

/*
Multiplexing Channel Inputs:

Write a program that starts several goroutines, each sending a sequence of numbers on its own channel.
Use a select statement to multiplex the channel inputs and process the numbers as they arrive from any goroutine.

*/

type MuxChan struct{}

func TestMultiChan(t *testing.T) {
	var mx = MuxChan{}
	mx.processSeqOfNum()
}
func (mx *MuxChan) processSeqOfNum() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			val := rand.Intn(20)
			ch1 <- val
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 2; i++ {
			val := rand.Intn(20)
			ch2 <- val
		}
		close(ch2)
	}()
	go func() {
		for i := 0; i < 2; i++ {
			val := rand.Intn(20)
			ch3 <- val
		}
		close(ch3)
	}()

	closedChannel := 0
	for {
		if closedChannel == 3 {
			break
		}
		select {
		case val, ok := <-ch1:
			if ok {
				fmt.Printf(" ch1:%d ", val)
			} else {
				fmt.Println("ch1 closed")
				closedChannel++
				ch1 = nil
			}
		case val, ok := <-ch2:
			if ok {
				fmt.Printf(" ch2:%d ", val)
			} else {
				fmt.Println("ch2 closed")
				closedChannel++
				ch2 = nil
			}

		case val, ok := <-ch3:
			if ok {
				fmt.Printf(" ch3:%d ", val)
			} else {
				fmt.Println("ch3 closed")
				closedChannel++
				ch3 = nil
			}
		}

	}
	fmt.Println("all done")
}

/*
Dining Philosophers Problem:

Solve the Dining Philosophers problem using goroutines and channels.
This classic concurrency problem involves a certain number of philosophers
who do nothing but think and eat, competing for a limited number of resources (forks).

*/

type DiningPhilosophers struct {
	fork         int
	philosophers int
}

func NewDiningPhilosophers(fork int, philosophers int) DiningPhilosophers {
	return DiningPhilosophers{fork: fork, philosophers: philosophers}
}
func TestDiningP(t *testing.T){
	forks:=5
	dining:=NewDiningPhilosophers(5,9)
	buffer:=make(chan int,forks)
	go dining.eat(buffer)
	go dining.think(buffer)

}
func(dp *DiningPhilosophers)eat(ch chan int){

}
func(dp *DiningPhilosophers)think(ch chan int){

}