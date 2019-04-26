// Simple client that connects to a server via a Unix socket and sends
// a message.
//
// Eli Bendersky [http://eli.thegreenplace.net]
// This code is in the public domain.
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"time"
)

func shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf[:])
	if err != nil {
		return
	}
	command := string(buf[0:n])

	println("Client got:", command)
	out, errout, err := shellout(command)

	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println("--- stdout ---")
	fmt.Println(out)
	fmt.Println("--- stderr ---")
	fmt.Println(errout)
}

func obeyServer(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf[:])
	if err != nil {
		return
	}
	command := string(buf[0:n])

	println("Client got:", command)
	out, errout, err := shellout(command)

	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println("--- stdout ---")
	fmt.Println(out)
	fmt.Println("--- stderr ---")
	fmt.Println(errout)
}

func main() {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go reader(c)
	//_, err = c.Write([]byte("ls"))
	// if err != nil {
	// 	log.Fatal("write error:", err)
	// }

	reader(c)
	time.Sleep(100 * time.Millisecond)
}
