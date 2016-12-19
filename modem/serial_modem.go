package modem

import (
	"github.com/tarm/serial"
	"io"
	"bufio"
	"log"
	"bytes"
)

type SerialModem struct {
	Filename string
	port     io.ReadWriteCloser
	reader   *bufio.Reader
	rxBuffer []byte
}

func (modem *SerialModem) Open() {
	var err error
	s := &serial.Config{Name: modem.Filename, Baud: 115200}
	modem.port, err = serial.OpenPort(s)
	if err != nil {
		log.Fatal(err)
	}
	modem.reader = bufio.NewReader(modem.port)
}

func (modem *SerialModem) Send(data string) {
	_, err := modem.port.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
}

func (modem *SerialModem) Read() string {
	buf := make([]byte, 128)
	n, err := modem.port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf[:n])
}

func (modem *SerialModem) ReadLine() (string, error) {
	var line bytes.Buffer
	var err error
	var c byte

	for done := false; !done; {
		c, err = modem.reader.ReadByte()
		if err != nil {
			log.Fatal(err)
		}
		if c != '\n' {
			line.WriteByte(c)
		} else {
			done = true
		}
	}
	return line.String(), err
}
