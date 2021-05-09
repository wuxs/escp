package escp

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
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
	datas := hex.EncodeToString(data)
	fmt.Println(datas)
	fmt.Println(data)
	return e.output.Write(data)
}

func (e *Escp) Print(str string) *Escp {
	e.buffer.Write([]byte(str))
	return e
}

// ESC/P
func (e *Escp) SelectJustification(align Align) *Escp {
	e.buffer.Write([]byte{0x1B, 0x61, byte(align)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SelectCharacterTable(char CharTable) *Escp {
	e.buffer.Write([]byte{0x1B, 0x74, byte(char)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SelectCharacterSet(c CharCode) *Escp {
	e.buffer.Write([]byte{0x1B, 0x52, byte(c)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SelectCharacterStyle(style CharStyle) *Escp {
	e.buffer.Write([]byte{0x1B, 0x71, byte(style)})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SelectBold() *Escp {
	e.buffer.Write([]byte{0x1B, 0x45})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) CancelBold() *Escp {
	e.buffer.Write([]byte{0x1B, 0x46})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SelectItalic() *Escp {
	e.buffer.Write([]byte{0x1B, 0x34})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) CancelItalic() *Escp {
	e.buffer.Write([]byte{0x1B, 0x35})
	return e
}

// ESC/P & ESC/P2
func (e *Escp) SelectDoubleStrike() *Escp {
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
func (e *Escp) BarCode(k BarType) *Escp {
	e.buffer.Write([]byte{0x1B, 0x28, 0x42, byte(k)})
	return e
}
