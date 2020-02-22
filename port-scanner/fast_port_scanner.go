package main

import (
	"fmt"
	"net"
	"sync"
)

//Run the command time -> give you some information about the execution time of the program
//Scanning a single port at a time isn’t useful, and it certainly isn’t efficient. TCP ports range from 1 to 65535;

const localhost = "127.0.0.1" //scan your ports

func main() {

	//Fast scanner
	//Using waitgroup will allows us to create connection in a new goroutine till the previous goroutine has been finished thread-safe
	var wg sync.WaitGroup
	for p := 1 ; p <= 65535 ; p++ {
		wg.Add(1)
		go func(host string, port int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d",host, port)
			conn,err := net.Dial("tcp",address)
			//CheckError(err)
			if err == nil {
				fmt.Printf("Host -> %s, The Connection has been established successfuly with port -> %d\n",host,port)
				conn.Close()
			}
		}(localhost,p)
	}

	wg.Wait()

}

//Once you’ve created WaitGroup, you can call a few methods on the struct. The first is Add(int),
//which increases an internal counter by the number provided. Next,
//Done() decrements the counter by one. Finally,
//Wait() blocks the execution of the goroutine in which it’s called,
//and will not allow further exe- cution until the internal counter reaches zero.
//You can combine these calls to ensure that the main goroutine waits for all connections to finish.

