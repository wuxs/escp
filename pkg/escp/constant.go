package escp

type Align byte

const (
	Left   Align = 0x0
	Center Align = 0x1
	Right  Align = 0x2
	Full   Align = 0x3
)

type CharTable byte

const (
	Italic      CharTable = 0x0
	PC437       CharTable = 0x1
	UserDefined CharTable = 0x2
)

type CharCode byte

const (
	USA      CharCode = 0x00
	France   CharCode = 0x01
	Germany  CharCode = 0x02
	UK       CharCode = 0x03
	Denmark1 CharCode = 0x04
	Sweden   CharCode = 0x05
	Italy    CharCode = 0x06
	Spain1   CharCode = 0x07
	Japan    CharCode = 0x08
	Norway   CharCode = 0x09
	Denmark2 CharCode = 0x10
	Spain2   CharCode = 0x11
	Latin    CharCode = 0x12
	Korea    CharCode = 0x13
	Legal    CharCode = 0x64
)

type CharStyle byte

const (
	None    CharStyle = 0x0
	Outline CharStyle = 0x1
	Shadow  CharStyle = 0x2
	Both    CharStyle = 0x3
)

type Status byte

const (
	OFF Status = 0x0
	ON  Status = 0x1
)

type BarType byte

const (
	EAN13   BarType = 0x00
	EAN8    BarType = 0x01
	ITF     BarType = 0x02
	UPCA    BarType = 0x03
	UPCE    BarType = 0x04
	CODE39  BarType = 0x05
	CODE128 BarType = 0x06
	POSTNET BarType = 0x07
)
