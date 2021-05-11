package main

import (
	"flag"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"github.com/wuxs/escp/pkg/escp"
)

func main() {
	dev := flag.String("d", "/dev/ttyS1", "serial dev name")
	options := serial.OpenOptions{
		PortName:        *dev,
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	s, err := serial.Open(options)
	if err != nil {
		fmt.Println(err.Error())
	}

	//s, _ := os.Create("test.txt") // write to file
	defer s.Close()

	esc := escp.NewEscP(s)
	esc.SpecifyAlignment(escp.Center).
		SpecifyBold().
		Print("Hello world\n")
	size, err := esc.Flush()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(size)
	}
}
