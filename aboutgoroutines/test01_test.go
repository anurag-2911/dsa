package aboutgoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestX(t *testing.T) {
	fmt.Println()
}

func TestBasicGoRoutine(t *testing.T) {
	go basicRoutine()
	time.Sleep(10 * time.Second)
	fmt.Println("all done")
}
func basicRoutine() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello ", i)
	}
}
func TestBasicChan(t *testing.T) {
	ch := make(chan int)
	go produce(ch)

	for val := range ch {
		fmt.Printf("%d\n", val)
	}

	// bufferred channel
	bch := make(chan int, 2)
	bch <- 1000
	bch <- 2000
	close(bch)
	for val := range bch {
		fmt.Printf("%v\t", val)
	}

}
func produce(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
func TestSelect(t *testing.T) {
	selectchannel()
}
func selectchannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "Hallo"
		close(ch1)

	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "hello"
		close(ch2)

	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

func square(ch chan int) {
	temp := <-ch
	ch <- temp * temp
	close(ch)
}
func TestSquare(t *testing.T) {
	ch := make(chan int)
	go square(ch)
	ch <- 9
	result := <-ch
	fmt.Println(result)
}

func TestMultiPSingleC(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go multiproducer(ch, &wg, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		fmt.Printf("%d\t", value)
	}

}
func multiproducer(ch chan<- int, wg *sync.WaitGroup, id int) {
	for i := 0; i < 5; i++ {
		ch <- i * id
	}
	wg.Done()
}

func TestSelectWTimeout(t *testing.T) {
	selectntimeout()
	fmt.Println("test completed")
}
func selectntimeout() {
	ch := make(chan string, 0)

	go func() {
		time.Sleep(10 * time.Second)
		ch <- "hallo"
	}()

	select {
	case result := <-ch:
		fmt.Printf("%s\n", result)
	case <-time.After(5 * time.Second):
		fmt.Printf("%s\n", "timeout")

	}
	fmt.Println("all done")
}



//worker pool pattern example
func TestWorkerPool(t *testing.T) {
	const numOfJobs = 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for w := 1; w <= 3; w++ {
		go worker(jobs, results, w)
	}

	for j := 1; j <= numOfJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for r := range results {
		fmt.Println(r)
	}

}

func worker(jobs <-chan int, results chan<- int, id int) {
	for job := range jobs {
		fmt.Printf("worker %d started job %d\n", id, job)
		time.Sleep(2 * time.Second)
		results <- job * 100
		fmt.Printf("worker %d completed job %d\n", id, job)
	}

}
