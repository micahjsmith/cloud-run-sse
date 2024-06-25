package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "text/event-stream")
		rw.Header().Add("Cache-Control", "no-cache")

		flusher := rw.(http.Flusher)

        rw.Write([]byte("data: a\n"))
		flusher.Flush()
		time.Sleep(time.Second * 1)
        rw.Write([]byte("data: b\n"))
		flusher.Flush()
		time.Sleep(time.Second * 1)
        rw.Write([]byte("data: c\n"))
		flusher.Flush()
		time.Sleep(time.Second * 1)
        rw.Write([]byte("data: d\n"))
		flusher.Flush()
		time.Sleep(time.Second * 1)
        rw.Write([]byte("data: e\n"))
        rw.Write([]byte("data: [DONE]\n"))
	})

	s := &http.Server{
		Addr: ":8080",
	}
	log.Fatal(s.ListenAndServe())
}
