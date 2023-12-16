package gobasics

import (
	"bufio"
	"encoding/json"
	"fmt"
	"sync"

	"net/http"
	"os"
	"testing"
	"time"
)

func TestX(t *testing.T) {
	fmt.Println()
}

func TestBasics(t *testing.T) {
	hallo()
}
func hallo() {
	fmt.Println("hallo Welt")
	var x int = 101
	var str1 string = "Hallo Welt"
	fmt.Printf("%d and %s\n", x, str1)
	// if else
	variable := 100
	if variable == 100 {
		fmt.Printf("%v\n", variable)
	} else {
		fmt.Printf("%v\n", "no match")
	}
	//loops
	for i := 0; i < 3; i++ {
		fmt.Printf("%d\t", i)
	}
	y := 10
	res := func(x int) int {
		fmt.Printf("%d\n", x*x)
		return x * x
	}(y)
	fmt.Printf("%d\t", res)
}
func TestArraysNSlices(t *testing.T) {
	arraysnslices()
}
func arraysnslices() {
	arr := [5]int{1, 2, 3, 4, 5}

	slc := arr[0:]
	slc = append(slc, 81)
	fmt.Println(arr, slc)
}
func TestMaps(t *testing.T) {
	mindmap()
}
func mindmap() {
	session := make(map[int]string, 0)
	session[1] = "one"
	session[2] = "two"
	for i, s := range session {
		fmt.Println(i, s)
	}
	fmt.Println(session)
}

func TestStructMethods(t *testing.T) {
	bundle := &Bundle{}
	bundle.Add("1", "one")
}

type Bundle struct {
	list map[string]string
}

func NewBundle() *Bundle {
	return &Bundle{list: map[string]string{}}
}
func (b *Bundle) Add(id string, Name string) {
	if b.list == nil {
		b.list = make(map[string]string)
	}
	b.list[id] = Name
}
func TestIFace(t *testing.T) {
	var action Action
	action = &FileCopy{}
	fmt.Println(action.actionName(), action.processAction())
	action = &InstallDir{}
	fmt.Println(action.actionName(), action.processAction())
}

type Action interface {
	processAction() bool
	actionName() string
}

type FileCopy struct {
}

func (fc *FileCopy) processAction() bool {
	time.Sleep(2 * time.Second)
	return true
}
func (fc *FileCopy) actionName() string {
	return "FileCopy"
}

type InstallDir struct {
}

func (id *InstallDir) processAction() bool {
	time.Sleep(time.Second * 2)
	return true
}
func (id *InstallDir) actionName() string {
	return "InstallDir"
}

func TestGoRnChan(t *testing.T) {
	ch := make(chan int)
	go processChan(ch)
	for r := range ch {
		fmt.Printf("%d\t", r)
	}
}
func processChan(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(5 * time.Second)
	}
	close(ch)
}

func TestErrorHandling(t *testing.T) {
	res, err := div(10, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	err = riskyfn()
	fmt.Println(err)
	fmt.Println(err.Error())

}
func div(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("can't divide by zero")
	}
	return x / y, nil
}

type XError struct {
	ErrorMessage string
}

func (xe *XError) Error() string {
	return xe.ErrorMessage
}
func riskyfn() error {
	return &XError{"kuch nahi ho sakta"}
}

func TestDefer(t *testing.T) {
	fmt.Println("opening the file")
	defer fmt.Println("closing the file")
	fmt.Println("processing the file")
	fmt.Println("some other actions")
	justsomework()
}
func justsomework() {
	time.Sleep(time.Second * 5)
}

func TestRangeOverChannel(t *testing.T) {
	ch := make(chan int, 0)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for c := range ch {
		fmt.Printf("%d\t", c)
	}

}

func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go OneS(ch1)
	go TwoS(ch2)
	select {
	case val := <-ch1:
		fmt.Println("by ", val)
	case val := <-ch2:
		fmt.Println("by ", val)
		// default:
		// 	println("default case")
	}

}
func OneS(ch chan int) {
	time.Sleep(4 * time.Second)
	ch <- 101
}
func TwoS(ch chan int) {
	time.Sleep(time.Second * 20)
	ch <- 4000
}

func TestHttpServer(t *testing.T) {
	http.HandleFunc("/ping", pingme)
	http.ListenAndServe(":8084", nil)
}
func pingme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hallo Welt")
}

func TestFileRead(t *testing.T) {
	file, err := os.Open("C:/test/test.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func TestWriteFile(t *testing.T) {
	file, err := os.Create("c:\\test\\two.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	n, err := file.WriteString("Hallo Welt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)

}

func TestJson(t *testing.T) {
	devicesession := &Session{ID: "999", Name: "device"}
	data, err := json.Marshal(devicesession)
	if err != nil {
		fmt.Println(err)
	}
	var ds Session
	err = json.Unmarshal(data, &ds)
	if err != nil {
		fmt.Println(err)
	}
}

type Session struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestGRnWaitGrp(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	go func() {
		for v := range ch {
			fmt.Printf("%d\t", v)
		}
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go dosometask(&wg, ch, i)
	}

	wg.Wait()
	close(ch)	

}
func dosometask(wg *sync.WaitGroup, ch chan int, i int) {
	defer wg.Done()
	// time.Sleep(5 * time.Second)
	ch <- i

}

func TestChannelSynnc(t *testing.T){
	ch:=make(chan bool)
	go sugnal(ch)
	fmt.Println(<-ch)

}
func sugnal(ch chan bool){
	time.Sleep(4*time.Second)
	ch<-true
}

