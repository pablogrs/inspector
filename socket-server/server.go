package main

import (
	"log"
	"net"
	"os"
	"time"

	slavecommandspb "github.hpe.com/pablo-gon-sanchez/inspector-gadget/protopb/commands"

	inspectorConfig "github.hpe.com/pablo-gon-sanchez/inspector-gadget/inspectorConfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// SockAddr using Unix socket
const SockAddr = "/tmp/echo.sock"

type server struct{}

func main() {
	inspectorConfig.LoadConfig()

	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("unix", SockAddr)

	if err != nil {
		log.Fatal("listen error:", err)
	}

	grpcServer := grpc.NewServer()
	// Register reflection service
	reflection.Register(grpcServer)

	slavecommandspb.RegisterCommandServiceServer(grpcServer, &server{})

	if grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v \n", err)
	}
}

func (*server) RunCommand(request *slavecommandspb.RunCommandRequest, stream slavecommandspb.CommandService_RunCommandServer) error {
	log.Printf("Client connected to RunCommand \n")

	for _, v := range inspectorConfig.InspectorConfiguration.Commands {
		//log.Printf("Commands: %v -%v \n", v.Name, v.Parameters)
		fullCommand := v.Name + " -" + v.Parameters

		time.Sleep(2000 * time.Millisecond)

		response := &slavecommandspb.RunCommandResponse{
			CommandResponse: fullCommand,
		}
		stream.SendMsg(response)
	}

	return nil
}
