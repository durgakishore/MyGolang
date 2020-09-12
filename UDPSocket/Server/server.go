package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	inputPath              string
	destIPAddress          string
	port                   string
	serverNetworkInterface string
	numOfPacketsPerSecond  int
	maxMessagesPerPacket   int
	symbol                 string
)

type packetData struct{}

func main() {

	/*[Server.exe -infile input.txt -multicast 225.1.1.7 -packetcount 3 --mmp 5 -symbol IBM 65071 192.168.195.4 ]
	  [    0 		  1       2          3         4       5          6    7  8    9     10    11      12]*/

	flag.StringVar(&inputPath, "infile", " ", "input file path")
	flag.StringVar(&destIPAddress, "multicast", " ", "Client IP Address")
	flag.StringVar(&symbol, "symbol", "ALL", "Message Symbol")

	flag.IntVar(&numOfPacketsPerSecond, "packetcount", 5, "packets per second")
	flag.IntVar(&maxMessagesPerPacket, "mmp", 5, "messages per packet")

	flag.Parse()

	cmdArguments := flag.Args()

	port = cmdArguments[0]
	serverNetworkInterface = cmdArguments[1]

	txtFileData := readInputFileData(inputPath)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	startSendingUDPPacket(txtFileData, wg)
	wg.Wait()
}

func readInputFileData(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	var fileData []string

	data, err := ioutil.ReadFile(fileName)

	fileText := string(data)
	str := strings.Split(fileText, "") // one packet upto "ETX"
	for _, msg := range str {
		if symbol == "ALL" || strings.Contains(msg, symbol) {
			fileData = append(fileData, msg)
		}
	}

	return fileData
}

func startSendingUDPPacket(txtFileData []string, wg *sync.WaitGroup) {

	defer wg.Done()
	address := serverNetworkInterface + ":" + port
	s, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 8196)

	for {
		_, addr, _ := connection.ReadFromUDP(buffer)
		//fmt.Printf("Client connected  %s\n", addr.String())
		go processClient(txtFileData, addr, connection)
	}
}

func processClient(txtFileData []string, addr *net.UDPAddr, connection *net.UDPConn) {
	/*number of packets per second handling timer tick
	numOfPacketsPerSecond -- configurable parameter 	*/

	duration := time.Tick(time.Duration((1000 / numOfPacketsPerSecond)) * time.Millisecond)

	var data []byte
	totalPacketsSent := 0
	nPacketsSentIndex := 0

	for {

		if len(txtFileData) <= nPacketsSentIndex {
			data = []byte("stop")
			_, err := connection.WriteToUDP(data, addr)
			if err != nil {
				fmt.Println(err)
			}
			break
		}

		<-duration

		if (len(txtFileData) - nPacketsSentIndex) < maxMessagesPerPacket {
			slice := txtFileData[nPacketsSentIndex:]
			data = []byte(strings.Join(slice, ","))
		} else {
			slice := txtFileData[nPacketsSentIndex : nPacketsSentIndex+maxMessagesPerPacket]
			data = []byte(strings.Join(slice, ","))
		}
		nPacketsSentIndex = nPacketsSentIndex + maxMessagesPerPacket

		//fmt.Println(time.Now())
		totalPacketsSent++
		//fmt.Printf("Sending : %s\n", string(data))
		_, err := connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println()
	fmt.Printf("*********Data sent to client %s********** \n", addr.String())
	fmt.Printf("Total  Messages in input files : %d\n", len(txtFileData))
	fmt.Printf("Max Messages per packet : %d\n", maxMessagesPerPacket)
	fmt.Printf("No. of packets per second : %d\n", numOfPacketsPerSecond)
	fmt.Printf("Total Packets sent: %d \n", totalPacketsSent)
	fmt.Printf("Total transmission time : %d seconds\n", totalPacketsSent/numOfPacketsPerSecond)
}
