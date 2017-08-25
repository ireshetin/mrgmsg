package App

import (
	"fmt"
	"log"
	"time"
	"net"
	"flag"
	"./Processes"
	"math/rand"
	"strconv"
)

func Run() (error) {
	return listenTCP()
}

func listenTCP() error {
	var listenAddr string

	flag.StringVar(&listenAddr, "listen-addr", "127.0.0.1:9001", "address to listen")
	flag.Parse()

	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		return fmt.Errorf("error listening on %q: %s", listenAddr, err)
	}
	defer closeListener(listener, err)

	fmt.Printf("Listening on %q\n", listenAddr)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("Error accepting connection %q: %s", listenAddr, err)
			time.Sleep(1000 * time.Millisecond)
			continue
		}

		go Processes.Run(conn, getSid())
	}
}

func getSid() string {
	// todo: переделать
	return strconv.Itoa(rand.Int())
}

func closeListener(listener net.Listener, err error) {
	closeErr := listener.Close()
	if closeErr != nil {
		log.Printf("can't close listen socket: %s", err)
	}
}
