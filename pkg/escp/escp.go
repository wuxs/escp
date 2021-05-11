package escp

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
)

type Escp struct {
	output io.Writer
	buffer bytes.Buffer
}

func NewEscP(output io.Writer) *Escp {
	return &Escp{
		output: output,
		buffer: bytes.Buffer{},
	}
}

func (e *Escp) Flush() (int, error) {
	data := e.buffer.Bytes()
	e.buffer.Reset()
	return e.output.Write(data)
}

func (e *Escp) Print(str string) *Escp {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GB18030.NewEncoder())
	b, _ := ioutil.ReadAll(reader)
	e.buffer.Write(b)
	return e
}

// ESC/P
func (e *Escp) SpecifyAlignment(align Align) *Escp {
	e.buffer.Write([]byte{0x1B, 0x61, byte(align)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SpecifyCharacterTable(char CharTable) *Escp {
	e.buffer.Write([]byte{0x1B, 0x74, byte(char)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SpecifyCharacterSet(c CharCode) *Escp {
	e.buffer.Write([]byte{0x1B, 0x52, byte(c)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SpecifyCharacterStyle(style CharStyle) *Escp {
	e.buffer.Write([]byte{0x1B, 0x71, byte(style)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SpecifyBold() *Escp {
	e.buffer.Write([]byte{0x1B, 0x45})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) CancelBold() *Escp {
	e.buffer.Write([]byte{0x1B, 0x46})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SpecifyItalic() *Escp {
	e.buffer.Write([]byte{0x1B, 0x34})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) CancelItalic() *Escp {
	e.buffer.Write([]byte{0x1B, 0x35})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SpecifyDoubleStrike() *Escp {
	e.buffer.Write([]byte{0x1B, 0x47})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) CancelDoubleStrike() *Escp {
	e.buffer.Write([]byte{0x1B, 0x48})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) Underline(status Status) *Escp {
	e.buffer.Write([]byte{0x1B, byte(status)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) DoubleWidth(status Status) *Escp {
	e.buffer.Write([]byte{0x1B, 0x57, byte(status)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) DoubleHeight(status Status) *Escp {
	e.buffer.Write([]byte{0x1B, 0x77, byte(status)})
	return e
}

// ESC/P & ESC/P2
// TODO BarCode
func (e *Escp) BarCode(k BarType) *Escp {
	e.buffer.Write([]byte{0x1B, 0x28, 0x42, byte(k)})
	return e
}
