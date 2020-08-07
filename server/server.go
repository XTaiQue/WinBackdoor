package server

import (
	"fmt"
	"io"
	"net"
	"os"
)

func runserver(protocol string, host string) error {
	listener, err := net.Listen(protocol, host)

	if err != nil {
		return err
	}

	fmt.Println("[+] Listening on ", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[+] Client not Accepted ! : ", err)
		} else {
			fmt.Println("[+] New Client Connected => ", conn.RemoteAddr())
			go ProcessClient(conn, os.Stdout) // Go Routine processClient
			go ProcessClient(os.Stdin, conn)  // Go Routine processClient
		}
	}

	return nil
}

func ProcessClient(sstream io.Reader, dstream io.Writer) {
	buffer := make([]byte, 1024) // make Slices of 1024 Bytes (Slices == vector in C++)
	for {
		var NumberBytes int
		var err error
		NumberBytes, err = sstream.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("[+] Read error : %s\n", err)
			}
			break
		}

		_, err = dstream.Write(buffer[0:NumberBytes]) // _ for not get value
		if err != nil {
			fmt.Printf("[+] Write Error : %s\n", err)
		}
	}
}

func Netcat(protocol string, host string) error {
	err := runserver(protocol, host)

	if err != nil {
		return err
	}

	return nil
}
