package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)
type indexHandler struct{
	content string
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"11111111111")

}

func pinffun(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}
func midfun(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "mid",
	})
}
func v1longin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "v1login",
	})
}
func v2longin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "v2login",
	})
}
func gin1() {
	r := gin.Default()
	r.Use(midfun)
	r.GET("/ping",pinffun)
	v1 := r.Group("/v1")
	{
		v1.GET("/login",v1longin)
	}

	// Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/login",v2longin)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func nativeHttp(){
	http.Handle("/2",&indexHandler{content:"3333333333"})
	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"22222")
	})

	server := http.Server{
		Handler:           &indexHandler{},
	}
	lis, _ := net.Listen("tcp", ":8081")
	server.Serve(lis)
}
func fileServer() {
	router := gin.Default()
	//router.Static("/", "H:/github/")
	router.StaticFS("/1", http.Dir("H:/github/ebook"))

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
func main() {
	//nativeHttp()
	//gin1()
	fileServer()


}
