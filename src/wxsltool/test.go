package main

import (
	"log"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "COM8", Baud: 115200, Parity: serial.ParityEven}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("你好"))

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", buf[:n])
}
