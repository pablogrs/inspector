// Simple client that connects to a server via a Unix socket and
// retrieves the messages of a grpc stream

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"time"

	slavecommandspb "github.hpe.com/pablo-gon-sanchez/inspector-gadget/protopb/commands"

	"google.golang.org/grpc"
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

func main() {
	addr := "/tmp/echo.sock"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}))

	if err != nil {
		log.Fatalf("Can't connect: %v", err)
	}

	defer conn.Close()
	c := slavecommandspb.NewCommandServiceClient(conn)

	if err != nil {
		log.Fatal(err)
	}

	getCommand(c)
}

// Get the commands the server has defined and execute them
func getCommand(c slavecommandspb.CommandServiceClient) {
	log.Printf("Start command response call \n")

	request := &slavecommandspb.RunCommandRequest{
		CommandToRun: 1,
	}
	c.RunCommand(context.Background(), request)
	streamResponse, err := c.RunCommand(context.Background(), request) 

	if err != nil {
		log.Fatalf("Error when calling command stream %v", err)
	}

	for {
		msg, err := streamResponse.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error when reading message %v ", err)
		}
		println("Client got:", msg.GetCommandResponse())
		out, errout, err := shellout(msg.GetCommandResponse())

		fmt.Println("--- stdout ---")
		fmt.Println(out)
		fmt.Println("--- stderr ---")
		fmt.Println(errout)
	}

}
