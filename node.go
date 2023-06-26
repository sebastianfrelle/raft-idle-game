package main

import (
	"log"
	"os"
	"strconv"

	"github.com/sebastianfrelle/raft/server"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected at least two arguments; got %d", len(os.Args))
	}
	portArg := os.Args[1]
	port, err := strconv.ParseUint(portArg, 10, 16)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(uint16(port))
	log.Printf("Listening on port %d\n", port)
	s.Start()
}

// func main() {
// 	portNo, err := portNumber()
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}

// 	address := fmt.Sprintf("%s:%d", HOST, portNo)

// 	listen, err := net.Listen(TYPE, address)
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}
// 	defer listen.Close()

// 	fmt.Printf("Listening on %s:%d\n", HOST, portNo)

// 	for {
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			log.Fatal(err)
// 			os.Exit(1)
// 		}

// 		go handleConnection(conn)
// 	}
// }

// func portNumber() (uint16, error) {
// 	requiredArgsNo := 2
// 	if len(os.Args) < requiredArgsNo {
// 		return 0, fmt.Errorf("Expected at least %d arguments, got %d", requiredArgsNo, len(os.Args))
// 	}

// 	portNo, err := strconv.ParseUint(os.Args[1], 10, 16)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return uint16(portNo), nil
// }

// func handleConnection(conn net.Conn) {
// 	buffer := make([]byte, 1024)
// 	_, err := conn.Read(buffer)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	time := time.Now().Format("2006-01-02T15:04:05Z07:00")

// 	_, err = conn.Write([]byte("Hi back!\n"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = conn.Write([]byte(time))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	conn.Close()
// }
