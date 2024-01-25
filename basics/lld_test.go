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
double linked list, elements are added to the front, when read they are deleted from the current position
and added to the front
when adding if capacity is reached last element is removed
*/
type XNode struct {
	next, prev *XNode
	key, value int
}
type XLRUCache struct {
	head, tail *XNode
	lrucache   map[int]*XNode
	capacity   int
}

func NewLRUCache(cap int) *XLRUCache {
	return &XLRUCache{
		lrucache: make(map[int]*XNode),
		capacity: cap,
	}
}
func (lru *XLRUCache) Add(newNode *XNode) {
	if lru.head == nil {
		lru.head = newNode
		lru.tail = newNode
		return
	}
	current := lru.head
	newNode.next = current

}

func TestLRUCache(t *testing.T) {
	lrucache := NewLRUCache(10)
	node1 := &XNode{
		key:   1,
		value: 101,
	}
	node2 := &XNode{
		key:   2,
		value: 102,
	}
	node3 := &XNode{
		key:   3,
		value: 103,
	}
	lrucache.Add(node1)
	lrucache.Add(node2)
	lrucache.Add(node3)
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
