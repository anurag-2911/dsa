package basics

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

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
func(shorturlsvc *XShortURlSvc)Resolve(shorturl string)(string,bool){
	shorturlsvc.mu.Lock()
	defer shorturlsvc.mu.Unlock()
	if origurl,found:=shorturlsvc.urlMap[shorturl];found{
		return origurl,found
	}
	return "",false
}
func TestShortURLSvc(t *testing.T){
	router:=gin.Default()
	shorturlsvc:=NewShortURlSvc("http://localhost:9090")
	router.GET("/:shorturl",func(c *gin.Context){
		param:=c.Param("shorturl")
		shorturl:=shorturlsvc.baseurl + "/"+param
		if origurl,found:=shorturlsvc.Resolve(shorturl);found{
			c.Redirect(http.StatusMovedPermanently,origurl)
		}else{
			c.String(http.StatusNotFound,"url not found")
		}

	})
	router.POST("/shorten",func (c *gin.Context){

		var shortreq struct{
			OriginalURL string `json:"originalurl"`
		}
		if c.BindJSON(&shortreq)==nil{
			shorturl:=shorturlsvc.Shorten(shortreq.OriginalURL)
			c.JSON(http.StatusOK,gin.H{"shorturl":shorturl})
		}else{
			c.JSON(http.StatusBadRequest,gin.H{"error ":"invalid request"})
		}
	})
	router.Run(":9090")
}

//end of shorturl basic implementation

// LRU cache



// end of LRU cache