package basic

import
(
	"testing"
	"fmt"
)

func TestXxxY(t *testing.T) {
	fmt.Println()
}

// demos new call
func TestNewOp(t *testing.T){
	var session *Session=new(Session)
	session.ID="123"
	session.Name="test"
	fmt.Println(*session)
	

}

type Session struct{
	ID string
	Name string
	Count int
	ch chan int
	mp map[int]string
}

