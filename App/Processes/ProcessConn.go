package Processes

import (
	"fmt"
	"log"
	"net"
	"time"

	"../Helper"
	//"context"
	//"bytes"
)

type tChanBufData struct {
	buf []byte
	//bufLen int
	err error
}

func Run(conn net.Conn, sid string) error {
	return handleConn(conn, sid)

}

func handleConn(conn net.Conn, sid string) error {
	defer func() {
		log.Println("Close connection #" + sid)
		conn.Close()
	}()

	buf := make([]byte, 1024)
	bufRead := make([]byte, 1024)
	promise := make(chan tChanBufData, 1)

	for {
		go func() {
			transferData(conn, []byte("#"+sid+": Enter command: "))
			reqLen, err := conn.Read(bufRead)

			buf, err = Helper.DecriptRSA(Helper.PrivateKey2048, string(bufRead[:reqLen]))

			if err != nil { // шифрования нет, костыль для разработки
				buf = bufRead[:reqLen]
				err = nil
				fmt.Println("No rsa =(")
			} else {
				fmt.Println("RSAAAAAAA-AAA-AAAaaaa =)")
			}

			fmt.Printf("'%s' \n", buf)

			promise <- tChanBufData{
				buf: []byte(buf),
				err: err,
			}
		}()

		select {
		case bufData := <-promise:
			// Получили данные

			if bufData.err != nil {
				return fmt.Errorf("can't read data from connection: %s", bufData.err)
			}

			err := transferData(conn, []byte("Message received. \n"))

			if err != nil {
				return err
			} else {
				log.Printf("#"+sid+": OK! got request of len %d bytes: %s", bufData.buf)
			}

			if err := processCommand(bufData, conn); err != nil {
				log.Println(err)
			}

			break
		case <-time.Tick(100 * time.Second):
			return nil
		}
	}

	return nil
}

// {"Command": "login", "Params": 1}

func transferData(conn net.Conn, msg []byte) error {
	if _, err := conn.Write(msg); err != nil {
		return fmt.Errorf("can't write to connection: %s", err)
	}

	return nil
}

// {"Command": "login", "Params": {"name":"vasya", "password":"qwerty"}}
