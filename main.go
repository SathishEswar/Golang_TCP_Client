package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
    if err != nil {
        log.Fatal("tcp server listener error:", err)
	}
	fmt.Println("server is listening...")
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err.Error())
        }

        go handleNewConnections(conn)
    }
}

func handleNewConnections(conn net.Conn) {
    bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')

    if err != nil {
        log.Println("client disconnected.")
        conn.Close()
        return
    }

    message := string(bufferBytes)
    Address := conn.RemoteAddr().String()
    response := fmt.Sprintf(message + " from " + Address + "\n")

    log.Println(response)

    conn.Write([]byte(response))

    handleNewConnections(conn)
}