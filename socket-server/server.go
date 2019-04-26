package main

import (
	"log"
	"net"
	"os"

	inspectorConfig "github.hpe.com/pablo-gon-sanchez/inspector-gadget/inspectorConfig"
)

// SockAddr using Unix socket
const SockAddr = "/tmp/echo.sock"

// func echoServer(c net.Conn) {
// 	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
// 	io.Copy(c, c)
// 	//exec.Command("/bin/sh", c)
// 	c.Close()
// }

// func echoServer(c net.Conn) {
// 	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
// 	defer c.Close()

// 	for {
// 		buf := make([]byte, 512)
// 		nr, err := c.Read(buf)
// 		if err != nil {
// 			return
// 		}

// 		data := buf[0:nr]
// 		println("Server got:", string(data))
// 		_, err = c.Write(data)
// 		if err != nil {
// 			log.Fatal("Write: ", err)
// 		}
// 	}
// }

func sendCommands(c net.Conn) {
	//var fullCommand string

	log.Printf("Inspector config: %v \n", inspectorConfig.InspectorConfiguration)

	for _, v := range inspectorConfig.InspectorConfiguration.Commands {
		log.Printf("Commands: %v -%v \n", v.Name, v.Parameters)
		fullCommand := v.Name + " -" + v.Parameters
		_, err := c.Write([]byte(fullCommand))
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}

	// for _, record := range inspectorConfig.InspectorConfiguration {
	// 	log.Printf("Commands: %s", record)

	// 	for _, value := range record.([]interface{}) {
	// 		//log.Printf(" type of %s ", reflect.TypeOf(value))
	// 		if commands, ok := value.(map[interface{}]interface{}); ok {

	// 			log.Printf("%v -%v", commands["name"], commands["parameters"])
	// 			command := commands["name"].(string)
	// 			parameters := commands["parameters"].(string)
	// 			fullCommand = command + " -" + parameters
	// 		}

	// 		// actually send command
	// 		_, err := c.Write([]byte(fullCommand))
	// 		if err != nil {
	// 			log.Fatal("Write: ", err)
	// 		}

	// 		fullCommand = ""
	// 	}
	// }
}

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	defer c.Close()

	// for {
	// 	buf := make([]byte, 512)
	// 	nr, err := c.Read(buf)
	// 	if err != nil {
	// 		return
	// 	}

	// data := buf[0:nr]
	// println("Server got:", string(data))
	sendCommands(c)
	//time.Sleep(5 * time.Second)
	// }
}

func main() {
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	inspectorConfig.LoadConfig()
	//fmt.Printf("Configuration %v \n", inspectorConfig.InspectorConfiguration)

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(conn)
	}

}
