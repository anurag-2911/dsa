package linkedlist

import
(
	"testing"
)
func TestAppend(t *testing.T) {
	node1:=&Node{Data: 100}
	linkedList:=&LinkedList{}
	linkedList.Append(node1.Data)
	linkedList.Append(200)
	linkedList.Append(300)
	linkedList.Append(400)
	linkedList.Append(500)
	linkedList.Append(600)
	linkedList.Traverse()
	linkedList.Remove(600)
	linkedList.Traverse()
}
