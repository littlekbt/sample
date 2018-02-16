package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"flag"
)

var http2 *bool

func main() {

	http2 = flag.Bool("p", false, "http2 mode")
	flag.Parse()
	
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/image.png", handleImage)
	err := http.ListenAndServeTLS(":18443", "conf/certs/cert_example.com.pem", "conf/certs/key_example.com.pem", nil)
	fmt.Println(err)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	var html string
	if *http2 {
	  pusher, ok := w.(http.Pusher)
    if ok {
      fmt.Println("Push /image.png")
      pusher.Push("/image.png", nil)
    }
		html = `<html><body><img src="/image.png" width="300" height="300"></body></html>`
	} else {
		w.Header().Add("Link","</go-proxy/image.png>; rel=preload; as=image")
		html = `<html><body><img src="/go-proxy/image.png" width="300" height="300"></body></html>`
	}

  w.Header().Add("Content-Type", "text/html")
  fmt.Fprintf(w, html)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	image, _ := ioutil.ReadFile("/go/image.png")
  w.Write(image)
}
