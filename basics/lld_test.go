package basics

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestXX(t *testing.T) {
	fmt.Println()
}

//short url basic implementation

type XShortURlSvc struct {
	mu      sync.Mutex
	urlMap  map[string]string
	baseurl string
}

func NewShortURlSvc(baseurl string) *XShortURlSvc {
	return &XShortURlSvc{
		urlMap:  make(map[string]string),
		baseurl: baseurl,
	}
}
func (shorturlsvc *XShortURlSvc) Shorten(originalurl string) string {
	shorturlsvc.mu.Lock()
	defer shorturlsvc.mu.Unlock()
	short := fmt.Sprintf("%s/%d", shorturlsvc.baseurl, len(shorturlsvc.urlMap)+1)
	shorturlsvc.urlMap[short] = originalurl
	return short
}
func (shorturlsvc *XShortURlSvc) Resolve(shorturl string) (string, bool) {
	shorturlsvc.mu.Lock()
	defer shorturlsvc.mu.Unlock()
	if origurl, found := shorturlsvc.urlMap[shorturl]; found {
		return origurl, found
	}
	return "", false
}
func TestShortURLSvc(t *testing.T) {
	router := gin.Default()
	shorturlsvc := NewShortURlSvc("http://localhost:9090")
	router.GET("/:shorturl", func(c *gin.Context) {
		param := c.Param("shorturl")
		shorturl := shorturlsvc.baseurl + "/" + param
		if origurl, found := shorturlsvc.Resolve(shorturl); found {
			c.Redirect(http.StatusMovedPermanently, origurl)
		} else {
			c.String(http.StatusNotFound, "url not found")
		}

	})
	router.POST("/shorten", func(c *gin.Context) {

		var shortreq struct {
			OriginalURL string `json:"originalurl"`
		}
		if c.BindJSON(&shortreq) == nil {
			shorturl := shorturlsvc.Shorten(shortreq.OriginalURL)
			c.JSON(http.StatusOK, gin.H{"shorturl": shorturl})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error ": "invalid request"})
		}
	})
	router.Run(":9090")
}

//end of shorturl basic implementation

// LRU cache

/*
double linked list, elements are added to the front,
when read they are deleted from the current position
and added to the front
when adding if capacity is reached last element is removed
*/

type LRUNode struct {
	key, value int
	prev, next *LRUNode
}

type LRUCache struct {
	capacity   int
	items      map[int]*LRUNode
	head, tail *LRUNode
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*LRUNode),
	}
}
func (lru *LRUCache) addToFront(node *LRUNode) {
	if lru.head == nil {
		lru.head = node
		lru.tail = node
		return
	}
	node.next = lru.head
	lru.head.prev = node
	lru.head = node
}
func (lru *LRUCache) deletelastNode() {
	if lru.tail == nil {
		return
	}
	prev := lru.tail.prev
	prev.next = nil
	lru.tail = prev
}
func (lru *LRUCache) deletespecificNode(node *LRUNode) {
	// if the head node
	if lru.head.key == node.key {
		node.prev = nil
		lru.head = node.next
		node.next.prev = nil
		return
	}
	// if the tail node
	if lru.tail.key == node.key {
		prev := lru.tail.prev
		prev.next = nil
		lru.tail = prev
		return
	}
	// rest of the cases
	current := lru.head
	for current != nil {
		if current.key == node.key {
			current.next.prev = current.prev
			current.prev.next = current.next
			return
		}
		current = current.next
	}

}
func (lru *LRUCache) traverse() {
	if lru.head == nil {
		return
	}
	current := lru.head
	for current != nil {
		fmt.Print(current.value)
		fmt.Print("  ")
		current = current.next
	}
	fmt.Println()
}
func (lru *LRUCache) Get(key int) (int, bool) {
	if node, found := lru.items[key]; found {
		lru.deletespecificNode(node)
		lru.addToFront(node)
		return node.value, true
	}
	return -1, false
}
func (lru *LRUCache) Put(key, value int) {
	if node, found := lru.items[key]; found {
		node.value = value
		lru.deletespecificNode(node)
		lru.addToFront(node)
		return
	}
	node := &LRUNode{
		key:   key,
		value: value,
	}
	lru.items[key] = node
	lru.addToFront(node)

	if len(lru.items) > lru.capacity {
		lru.deletelastNode()
		delete(lru.items, key)

	}

}
func TestLRUCache(t *testing.T) {
	lrucache := NewLRUCache(5)
	lrucache.Put(101, 1001)
	lrucache.Put(102, 1002)
	lrucache.Put(103, 1003)
	lrucache.Put(104, 1004)
	lrucache.Put(105, 1005)
	lrucache.Put(106, 1006)
	lrucache.traverse()
	// lru.deletelastNode()
	// newFunction()
}

func newFunction() {
	node1 := &LRUNode{key: 101, value: 1001}
	node2 := &LRUNode{key: 102, value: 1002}
	node3 := &LRUNode{key: 103, value: 1003}
	node4 := &LRUNode{key: 104, value: 1004}
	node5 := &LRUNode{key: 105, value: 1005}
	node6 := &LRUNode{key: 106, value: 1006}

	lru := NewLRUCache(5)
	lru.addToFront(node1)
	lru.addToFront(node2)
	lru.addToFront(node3)
	lru.addToFront(node4)
	lru.addToFront(node5)
	lru.addToFront(node6)
	lru.traverse()

	lru.traverse()
	lru.deletespecificNode(node4)
	lru.traverse()
}

// end of LRU cache

/* Rate Limiter

Token bucket algorithm.
This algorithm allows for a certain number of requests (tokens) to be processed in a given time frame,
replenishing the tokens at a steady rate.

end of Rate Limiter*/

type RateLimiter struct {
	interval time.Duration
	maxToken int
	tokens   int
	lock     sync.Mutex
	ticker   time.Ticker
}

func NewRateLimiter(maxtokens int, interval time.Duration) *RateLimiter {
	rl := &RateLimiter{
		interval: interval,
		maxToken: maxtokens,
		tokens:   maxtokens,
		ticker:   *time.NewTicker(interval),
	}
	go rl.start()
	return rl
}
func (rl *RateLimiter) start() {
	for range rl.ticker.C {
		rl.lock.Lock()
		defer rl.lock.Unlock()
		if rl.tokens < rl.maxToken {
			rl.tokens++
		}
	}
}
func (rl *RateLimiter) allow() bool {
	rl.lock.Lock()
	defer rl.lock.Unlock()

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}
func TestRateLimiter(t *testing.T) {
	// Create a rate limiter that allows 5 requests per second
	limiter := NewRateLimiter(5, time.Second)

	// Example usage
	for i := 0; i < 10; i++ {
		if limiter.allow() {
			fmt.Println("Request", i, "allowed")
		} else {
			fmt.Println("Request", i, "denied")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

//end of rate limiter

/*
Basic Rate limiter using a ticker
*/

type RateL struct{}

func (rl *RateL) ratelimit() {
	ratelimiter := time.NewTicker(1 * time.Second)
	defer ratelimiter.Stop()
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ratelimiter.C:
				{
					doSomeAction()
				}
			case <-done:
				{
					fmt.Println("exiting rate limiter")
					return
				}
			}
		}
	}()
	time.Sleep(10*time.Second)
	done <- true
}
func doSomeAction() {
	time.Sleep(time.Second)
	fmt.Println("done at ", time.Now())
}
func TestRateLimit(t *testing.T) {
	rl := &RateL{}
	rl.ratelimit()
}
