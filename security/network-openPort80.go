
// this shows only hosts with open port 80

package main

import (
    "fmt"
    "net"
    "sync"
	"time"
)

func main() {
    var wg sync.WaitGroup
    ipBase := "192.168.178.0"

    for i := 1; i < 255; i++ {
        ip := ipBase + fmt.Sprint(i)
        wg.Add(1)
        go func(ip string) {
            defer wg.Done()
            if _, err := net.DialTimeout("tcp", ip+":80", 1000*time.Millisecond); err == nil {
                fmt.Println(ip, "is up")
            }
        }(ip)
    }

    wg.Wait()
}
