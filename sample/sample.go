package main

import(
  "time"
  "fmt"
	"crypto/rand"
  "runtime"
  "github.com/samarudge/go-stats"
)

func main() {
  gostats.Start("localhost:8127", 1, "sample")
  time.Sleep(5*time.Second)

  var randBytes []byte
  bytesToRead := 128*1024
  for i := 0;  i<=bytesToRead; i++ {
    byteChunk := make([]byte, 1024)
    _, err := rand.Read(byteChunk)
    if err != nil {
      fmt.Println("error:", err)
      return
    }

    randBytes = append(randBytes, byteChunk...)

    time.Sleep((60*time.Second)/time.Duration(bytesToRead))
  }

  time.Sleep(30)

  randBytes = []byte{}
  runtime.GC()

  time.Sleep(10)
}
