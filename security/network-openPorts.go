
// this shows only hosts with open ports <= 1000

package main

import (
	"fmt"
	"net"
	"sync"
	"time"
	"strconv"
	"sort"
)


func main() {
	lines := []string{}
	var hostList[256]string 
    var wg sync.WaitGroup
    ipBase := "192.168.178."
	j:=1
    for i := 1; i < 255; i++ {
        ip := ipBase + fmt.Sprint(i)
        wg.Add(1)
        go func(ip string) {
            defer wg.Done()
			for port :=1; port < 1023; port ++ {
				if _, err := net.DialTimeout("tcp", ip + ":" + strconv.Itoa(port), 20*time.Millisecond); err == nil {
					//fmt.Println(ip, strconv.Itoa(port) + " is up")
					hostList[j] = ip + " " + strconv.Itoa(port) + " is up"
					lines = append(lines, ip + "\t" + strconv.Itoa(port) + "\t is up") 
				}
			}
        }(ip)
		j++
    }
    wg.Wait()

	sort.Strings(lines)
	for _, element := range lines {
		fmt.Println(element)
	}


	fmt.Println("done") 
}



