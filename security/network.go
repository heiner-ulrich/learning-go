package main
// monitoring of hosts in subnet. 
//   Shows also hosts with no open ports.
//   gets subnet dynamically

// todos:
// ping alternative nmap 192.168.178.0/24 -sn
// subnet as parameter on startup or evaluate at startup (only /24 subnet) - DONE
// loop / service
// writing to logfile
// updating status file or database (asset and network inventory)
// notifying on new hosts in subnet
// new hosts: nmap -A -T4 -Pn
// notifying on hosts down (uc-projects)
// try that on raspi
// nmap own public ip or own portscanner: https://medium.com/@KentGruber/building-a-high-performance-port-scanner-with-golang-9976181ec39d

// compare nmap 192.168.178.0/24


import (
	"fmt"
	"sync"
	"strconv"
	"os/exec"
	"runtime"	
	"time"
	"regexp"
	"net"
	"log"
)


func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		printTime("\nstart execution: ")
		subnet_ping()
		getPublicIP()
	}
}



// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println("\nLocal IP of this machine: ", localAddr.IP)
    return localAddr.IP
    
}




func printTime(msg string) { 
	fmt.Println(msg, time.Now().Format("2006-01-02 15:04:05")) 
} 



func getPublicIP () {
	out, err := exec.Command("curl", "ipinfo.io/ip").Output()
	if err != nil {		
		fmt.Println("Fehler getPublicIP : \n", err)
		
	}
	fmt.Println("\nPublic IP of this network ", string(out[:])) 	
}



func subnet_ping() {
	// subnet:
	// before:  var subnet = "192.168.178."
	subnet := GetOutboundIP().String()	
	regex, _ := regexp.Compile(`^[\d^\.]+\.`) 
	subnet = regex.FindString(subnet)
	fmt.Println("\nsubnet: ", subnet)	
	// async:
	var wg sync.WaitGroup 
	wg.Add(256)
	// reporting:
	var hostList [256][5]string //0-ip 1-pingstatus 2-pingresult 3-hostname 4-time
	hostsUp :=0	

	// async ping of subnet:
	for i := 0; i < 256; i++ {
		go func(i int) {
			defer wg.Done()
			ip2ping := subnet + strconv.Itoa(i)			
			m_pingstatus, m_pingresult, m_hostname, m_exectime := pinger(ip2ping)
			hostList[i][0] = ip2ping
			hostList[i][1] = m_pingstatus
			hostList[i][2] = m_pingresult
			hostList[i][3] = m_hostname
			hostList[i][4] = m_exectime
		}(i)
	}
	wg.Wait()
	
	// count hosts up and report:	 	
	for i := 0; i < 256; i++ {
		if hostList[i][1] == "OK" {
			hostsUp++
			fmt.Println(hostList[i][0]+"\t"+hostList[i][1]+"\t"+hostList[i][3]+"\t\t"+hostList[i][2]+"\t"+hostList[i][4]+"\t")  
		}
	}
	fmt.Println("\nHosts up: "+strconv.Itoa(hostsUp)+"\nFinished for subnet " + subnet + " Time: " +time.Now().Format("2006-01-02 15:04:05")  + "\n")
}





func pinger(ip2ping string)(string, string, string, string) {	
	exectime := time.Now().Format("2006-01-02 15:04:05") 
	hostname := "NO-HOST"
	pingstatus := "NO-PING" 
	pingresult := "NO-PINGRESULT"
	
	// ping ip in subnet:
	out, err := exec.Command("ping", ip2ping, "-t3").Output()
	if err != nil {		
		return pingstatus, pingresult, hostname, exectime
	}
	pingstatus = "OK"
	myRegexPing, _ := regexp.Compile(`min.+ms`)
	pingresult = myRegexPing.FindString(string(out[:]))
	//	fmt.Println(string(out[:])) // testing

	
	// get hostname:
	host, err_host := exec.Command("host", ip2ping).Output()
	//	fmt.Println("getting host: ", ip2ping) // testing
	if err_host != nil {		
		myRegex, _ := regexp.Compile(`[^\.]+$`) //get back i
		match := myRegex.FindString(string (ip2ping[:]))
		switch match {
			case "0":
			hostname = " (.0: network identifier?)"
			case "255": 
			hostname = " (.255: broadcast?)"	
			default:
			hostname = " (DNS lookup fail)"
		}		
		return pingstatus, pingresult, hostname, exectime
	}	
	
	// clean up and return values:
	myRegex, _ := regexp.Compile(`name pointer (.*)`)
	hostname = myRegex.FindString(string (host[:]))
	return pingstatus, pingresult, hostname, exectime
	
}




