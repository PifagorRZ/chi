package src

import (
	"fmt"
	"net"
)

func Send(s string) {
	fmt.Println(s)
}

func GetTCPSender(c net.Conn) func(string) error {

	return func(s string) error {
		// defer c.Close()
		fmt.Printf("Send: %s\n", s)
		_, err := fmt.Fprintf(c, s)

		return err

	}

}

type Sender = func(string) error