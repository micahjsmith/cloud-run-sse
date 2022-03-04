package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Length", "10")

		flusher := rw.(http.Flusher)

		rw.Write([]byte("a\n"))
		flusher.Flush()
		time.Sleep(time.Second * 1)
		rw.Write([]byte("b\n"))
		flusher.Flush()
		time.Sleep(time.Second * 1)
		rw.Write([]byte("c\n"))
		flusher.Flush()
		time.Sleep(time.Second * 1)
		rw.Write([]byte("d\n"))
		flusher.Flush()
		time.Sleep(time.Second * 1)
		rw.Write([]byte("e\n"))
	})

	s := &http.Server{
		Addr: ":8080",
	}
	log.Fatal(s.ListenAndServe())
}
