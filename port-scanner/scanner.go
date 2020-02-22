package main

import (
	"fmt"
	"log"
	"net"
)

//Run the command time -> give you some information about the execution time of the program
//Scanning a single port at a time isn’t useful, and it certainly isn’t efficient. TCP ports range from 1 to 65535;

const host = "scanme.nmap.org" //scan your ports

func main() {

	for p := 1 ; p <= 1024 ; p++ {
		CheckConnection(host,p)
	}

}

func CheckConnection(host string, port int) {
	address := fmt.Sprintf("%s:%d",host, port)
	conn,err := net.Dial("tcp",address)
	CheckError(err)
	if err == nil {
		fmt.Printf("Host -> %s, The Connection has been established successfuly with port -> %d\n",host,port)
		conn.Close()
	}

}

func CheckError(err error) {
	if err != nil {
		log.Printf("The connection doesn't get established well !!",err)
	}
}
