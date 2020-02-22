package main

import (
	"fmt"
	"net"
	"sync"
)

//The worker(int, *sync.WaitGroup) function takes two arguments: a channel of type int and a pointer to a WaitGroup.
//The channel will be used to receive work, and the WaitGroup will be used to track when a single work item has been completed.

const localhost_ = "127.0.0.1" //scan your ports

//This is a slight performance increase, as it will allow all the workers to start immediately.
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		CheckConn(localhost_,p)
		wg.Done()
	}
}

func CheckConn(host string, port int) {
	address := fmt.Sprintf("%s:%d",host, port)
	conn,err := net.Dial("tcp",address)
	if err == nil {
		fmt.Printf("Host -> %s, The Connection has been established successfuly with port -> %d\n",host,port)
		conn.Close()
	}

}

func main() {
	//Buffered channels are ideal for maintaining and tracking work for multiple producers and consumers. You’ve capped the channel
	//at 100, meaning it can hold 100 items before the sender will block.
	ports := make(chan int,100) //100 -> buffred channel that mean i can receive before i send the value inside the channel because an unbuffred channel have just one empty zone
	wg := sync.WaitGroup{}

	for i := 0; i< cap(ports); i++ {
		go worker(ports,&wg)
	}

	for i := 1; i<65535 ; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Done()
	close(ports)
}