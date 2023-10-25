package hashtable

import 
(
	"fmt"
	"sync"
)

func TestMain() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic ", r)
		}
	}()
	testfunctions()
}
func testfunctions() {
	fmt.Println("testing a generic hash map")
}

// implementing a generic hashtable

const BucketSize = 20

type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}
type GHashTable[K comparable, V any] struct {
	Bucket [BucketSize][]KeyValuePair[K, V]
	mu sync.Mutex
}

func (this *GHashTable[K, V]) hash(key K) int {
	sum := 0
	for _, val := range fmt.Sprintf("%v", key) {
		sum += int(val)
	}
	return sum
}
func (this *GHashTable[K, V]) Put(key K, value V) {
	this.mu.Lock()
	defer this.mu.Unlock()
	index := this.hash(key)
	bucket := &this.Bucket[index]

	for i, bk := range *bucket {
		if bk.Key == key {
			(*bucket)[i].Value = value
			return
		}
	}
	(*bucket) = append((*bucket), KeyValuePair[K, V]{Key: key, Value: value})
}

func (this *GHashTable[K, V]) Get(key K) (V, bool) {
	this.mu.Lock()
	defer this.mu.Unlock()

	index := this.hash(key)
	bucket := &this.Bucket[index]
	for _, val := range *bucket {
		if key == val.Key {
			return val.Value, true
		}
	}
	var zeroValue V
	return zeroValue, false
}

func (this *GHashTable[K, V]) Delete(key K) {
	this.mu.Lock()
	defer this.mu.Unlock()
	
	index := this.hash(key)
	bucket := &this.Bucket[index]
	for i, k := range *bucket {
		if k.Key == key {
			// Using a more explicit way to exclude the element
			*bucket = append((*bucket)[:i], (*bucket)[i+1:]...)
			return
		}
	}
}
