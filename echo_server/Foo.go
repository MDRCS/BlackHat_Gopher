package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct {}

func (fooReader *FooReader) Read(msg []byte) (int,error) {
	fmt.Print("in>")
	return os.Stdin.Read(msg)
}

type FooWriter struct {}

func (fooWriter *FooWriter) Write(msg []byte) (int,error) {
	fmt.Print("out>")
	return os.Stdout.Write(msg)
}

func main() {

	var (
		reader FooReader
		writer FooWriter
	)

	//input := make([]byte,4096)
	//
	//res,err := reader.Read(input)
	//if err != nil {
	//	log.Fatalln("Unable to write data")
	//}
	//
	//fmt.Printf("Read %d bytes from stdin\n", res)
	//
	//res,err = writer.Write(input)
	//if err != nil {
	//	log.Fatalln("Unable to write data")
	//}
	//
	//fmt.Printf("Wrote %d bytes to stdout\n", res)


	//Alternative code to acheive thne same result
	res, err := io.Copy(&writer, &reader)
	if err != nil {
		log.Fatalln("Unable to read/write data")
		os.Exit(1)
	}

	fmt.Println("Wrote %d bytes to stdout\n", res)


}