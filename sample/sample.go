package main

import (
	"crypto/rand"
	"fmt"
	"github.com/samarudge/go-stats"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	go http.ListenAndServe("localhost:6060", nil)

	gostats.Start("localhost:8127", 1, "sample")

	time.Sleep(5 * time.Second)

	fmt.Println("Starting generation of random bytes")
	var randBytes []byte
	bytesToRead := 128 * 1024
	for i := 0; i <= bytesToRead; i++ {
		byteChunk := make([]byte, 1024)
		_, err := rand.Read(byteChunk)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		randBytes = append(randBytes, byteChunk...)

		time.Sleep((60 * time.Second) / time.Duration(bytesToRead))
	}
	fmt.Println(" - Done")
	time.Sleep(30)

	fmt.Println("Forcing GC")
	randBytes = []byte{}
	runtime.GC()
	fmt.Println("Done")

	time.Sleep(10)
}
