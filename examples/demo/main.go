package main

import (
	"fmt"
	"github.com/wuxs/escp/pkg/escp"
	"os"
)

func main() {

	//options := serial.OpenOptions{
	//	PortName:        "/dev/ttyS2",
	//	BaudRate:        9600,
	//	DataBits:        8,
	//	StopBits:        1,
	//	MinimumReadSize: 4,
	//}
	//s, err := serial.Open(options)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	s, _ := os.Create("test.txt")  // write to file
	defer s.Close()
	e := escp.NewEscP(s)
	e.Print("Hello 江苏交控\n")
	e.SelectBold()
	e.Print("Hello\n")
	e.SelectItalic()
	e.Print("Hello 江苏交控\n")
	e.CancelBold()
	e.CancelItalic()

	e.Underline(escp.ON)
	e.Print("Hello\n")
	e.Print("Hello 江苏交控\n")
	e.Underline(escp.OFF)

	e.DoubleWidth(escp.ON)
	e.DoubleHeight(escp.ON)
	e.Print("Hello\n")
	e.Print("Hello 江苏交控\n")
	e.DoubleWidth(escp.OFF)
	e.DoubleHeight(escp.OFF)

	_, err := e.Flush()
	if err != nil {
		fmt.Println(err.Error())
	}
}
