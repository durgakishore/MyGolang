package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gosuri/uilive"
	"github.com/hidez8891/shm"
)

var (
	ipAddress              string
	port                   string
	serverNetworkInterface string
)

const (
	symbolIBM = "IBM"
	symbolBAC = "BAC"
)

func main() {

	//-multicast 225.1.1.7 65071 192.168.195.126

	flag.StringVar(&ipAddress, "multicast", " ", "Local IP address")

	flag.Parse()

	cmdArguments := flag.Args()

	port = cmdArguments[0]
	serverNetworkInterface = cmdArguments[1]

	wg := &sync.WaitGroup{}
	wg.Add(1)
	startUDPClientProcessing(wg)
	wg.Wait()
}

func startUDPClientProcessing(wg *sync.WaitGroup) {
	address := serverNetworkInterface + ":" + port
	s, err := net.ResolveUDPAddr("udp4", address)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	defer c.Close()

	data := []byte("trying to send a sample request to server\n")
	_, err = c.Write(data)

	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	fmt.Printf("\nUDP server is connected and server address is %s\n\n", c.RemoteAddr().String())
	totalPacketsReceived := 0

	var key, value string

	var IBMValue, BACVaule string

	writer := uilive.New() // updating terminal output
	writer.Start()

	w, _ := shm.Create("shm_name", 8196)
	defer w.Close()

	r, _ := shm.Open("shm_name", 8196)
	defer r.Close()

	for {
		buffer := make([]byte, 8196) //supported upto 35 messages in one packet
		_, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			return
		}

		w.WriteAt(buffer, 0)

		rbuf := make([]byte, 8196)

		r.ReadAt(rbuf, 0)

		if string(rbuf[0:4]) == "stop" {
			fmt.Println("\nTotal packets received : ", totalPacketsReceived)
			writer.Stop()
			wg.Done()
			break
		}

		totalPacketsReceived++
		//fmt.Printf("\nReceived Data : %s\n", string(rbuf))

		rtemp := bytes.Split(rbuf, []byte(","))

		for v := range rtemp {

			resultbuf := rtemp[v]
			n := bytes.Index(resultbuf, []byte("|5="))
			m := bytes.Index(resultbuf, []byte("|8="))
			k := bytes.Index(resultbuf, []byte("|9="))

			if n != -1 && m != -1 && k != -1 {
				a := n + 3
				b := a + (m - n - 3)

				c := m + 3
				d := c + (k - m - 3)

				key = string(resultbuf[a:b])
				value = string(resultbuf[c:d])

				if key == symbolIBM {
					IBMValue = value
				} else if key == symbolBAC {
					BACVaule = value
				}

				fmt.Fprintf(writer, "%s : %s \n%s : %s\n", symbolIBM, IBMValue, symbolBAC, BACVaule)
				time.Sleep(time.Millisecond * 5)
			}
		}

	}
}
