package main

import(
	"net/http"
	"os"
	"strconv"
	"fmt"
	"github.com/astaxie/beego/grace"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.Form.Get("action"))
	w.Write([]byte("WORLD!"))
	w.Write([]byte("ospid:" + strconv.Itoa(os.Getpid())))
}

type HttpHandler struct {
	pre func(http.ResponseWriter, *http.Request) error
	h func(http.ResponseWriter, *http.Request)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/*", handler)
	//
	err := grace.ListenAndServe("localhost:8081", mux)
	if err != nil {
		log.Println(err)
	}
	//log.Println("Server on 80801 stopped")
	//os.Exit(0)

	//http.ListenAndServe("localhost:8081", mux)
	//h := new(HttpHandler)
	//h.h = handler
	//mux.Handle("hello", h)
	//go func() {
	//	mux.Handle()
	//}
}
