package main

import (
	"fmt"
	"net"
	"sort"
)

//The worker(int, *sync.WaitGroup) function takes two arguments: a channel of type int and a pointer to a WaitGroup.
//The channel will be used to receive work, and the WaitGroup will be used to track when a single work item has been completed.

const home = "127.0.0.1" //scan your ports

//This is a slight performance increase, as it will allow all the workers to start immediately.
func Worker(ports,result chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d",home, p)
		conn,err := net.Dial("tcp",address)

		if err != nil {
			result <- 0
			continue
		}

		conn.Close()
		result <- p
	}
}


func main() {
	//Buffered channels are ideal for maintaining and tracking work for multiple producers and consumers. You’ve capped the channel
	//at 100, meaning it can hold 100 items before the sender will block.
	ports := make(chan int,100) //100 -> buffred channel that mean i can receive before i send the value inside the channel because an unbuffred channel have just one empty zone
	results := make(chan int) // unbuffred channel
	openports := []int{}

	for i := 0; i< cap(ports); i++ {
		go Worker(ports,results)
	}

	go func() {
		for i := 1; i<1024 ; i++ {
			ports <- i
		}
	}()

	for i := 1;i<1024 ; i++ {
		port := <- results
		if port!=0 {
			openports = append(openports,port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)

	for _,port := range openports{
		fmt.Printf("Port : %d is Open \n",port)
	}
}