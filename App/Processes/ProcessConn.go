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
			//QQLJK65zVhLdTde305P3yKkQh2kMCMH1EWfx5k6swzaoIg9CCH/WmTY72c4qNbBQxwDaVVkHVDZVQ9unvdsQlLaw8banOyhcLScr7F920k04bsyVGKKWsib6amY+1kaSbzQFnuWCPpZRkNW4PCQak7hSbdzCbW90lDN57PaJITfiQaqtyQmiP34572SLLymTypQj5RhCiLPx2Wv/8YiOChfFJqU+eBSHxjluVZOMjUOdixMfGct8Mt4KQQ0bEuUcyWE8hd8vCKNFFptq13RyUssyHhHNrjD2EiSPeLur/XffgAP2x0xX1Pe4agutx7pgWadotB8ULNaBmpsqOGbbwg==

			buf, err = Helper.DecriptRSA(Helper.PrivateKey2048, string(bufRead[:reqLen]))

			if err != nil { // шифрования нет, костыль для разработки
				buf = bufRead[:reqLen]
				err = nil
				fmt.Println("No rsa =(")
			} else {
				fmt.Println("RSAAAAAAA-AAA-AAAaaaa =)")
			}

			fmt.Printf("'%s'", buf)

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

// {"Command": "login", "Data": 1}

func transferData(conn net.Conn, msg []byte) error {
	if _, err := conn.Write(msg); err != nil {
		return fmt.Errorf("can't write to connection: %s", err)
	}

	return nil
}

// {"Command": "login", "Data": {"name":"vasya", "password":"qwerty"}}
