package server

import (
	"fmt"
	"net"
)

const HOST = "localhost"
const NETWORK = "tcp"

type server struct {
	network string
	address string
}

// TODO: to be altered later
type state uint64

/*
Do some bit manipulation probably.
*/

const MASK = 0b11111000

type verb uint8

// TODO: Make `go generate` just to try it out.
//go:generate stringer -type verb

const (
	// Just to pick something random
	POST verb = 0b110
)

const (
	VERB_LENGTH           = 3
	PAYLOAD_LENGTH_LENGTH = 8
)

func New(port uint16) *server {
	return &server{
		network: NETWORK,
		address: fmt.Sprintf("%s:%d", HOST, port),
	}
}

func (s *server) Start() error {
	listener, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	defer listener.Close()

	// Initialize the state here
	// var st state = 0

	for {
		_, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("Accepting message failed with %w\n", err)
		}

		// Read incoming message.
		// Let's just say that there's some internal state that gets modified.
		// Let's just say that there's exactly one type of message that can be sent from server
		// to server to keep things simple for now.
		// handleIncomingMessage(&conn, &st)
	}
}

type message struct {
	Verb    verb
	Payload string
}

func parseMessage(content []byte) (msg *message, err error) {
	// parse verb
	// parse length
	// parse payload

	verb, err := parseVerb(content)
	if err != nil {
		return msg, fmt.Errorf("parsing verb failed with %w", err)
	}

	// Parse length
	payloadLength, err := parsePayloadLength(content)
	if err != nil {
		return msg, fmt.Errorf("parsing payload length failed with %w", err)
	}

	// Read payload
	// TODO: consider passing in the offset: const offset = VERB_LENGTH + PAYLOAD_LENGTH_LENGTH
	payload := parsePayload(content, payloadLength)

	return &message{verb, payload}, nil
}

func parseVerb(content []byte) (verb, error) {
	// Get the first byte
	if len(content) < 1 {
		return 0, fmt.Errorf("Expected content of at least length 1; found length %d", len(content))
	}

	b := content[0]
	// Shift to the right to remove the portion of the number that doesn't denote the verb.
	// e.g. 0b11011011 becomes 0b00000110 (verb: POST).
	v := b >> (8 - VERB_LENGTH)
	switch v {
	case uint8(POST):
		return POST, nil
	default:
		return 0, fmt.Errorf("Found unrecognized verb %b", v)
	}
}

func parsePayloadLength(content []byte) (uint8, error) {
	if len(content) < 2 {
		return 0, fmt.Errorf("Expected content of at least length 2; found length %d", len(content))
	}

	// First number (remove the part denoting the verb)
	firstPart := content[0] << VERB_LENGTH
	// Second number (remove the part not denoting the payload length)
	secondPart := content[1] >> (PAYLOAD_LENGTH_LENGTH - VERB_LENGTH)

	// XOR the two numbers to yield the payload length
	payloadLength := firstPart ^ secondPart

	return payloadLength, nil
}

func parsePayload(content []byte, payloadLength uint8) string {
	// TODO: find a way of passing in the offset to bitshift by.
	const offset = (VERB_LENGTH + PAYLOAD_LENGTH_LENGTH) / 8
	buf := make([]byte, payloadLength)
	for i := 0; i < len(buf); i++ {
		/*
			Quick explainer:
			- To get the payload out, we
			- leftshift the byte to the left by 8 - (PAYLOAD_LENGTH_LENGTH - VERB_LENGTH) = 3
			- take the next byte and rightshift it by 8 - (8 - (PAYLOAD_LENGTH_LENGTH - VERB_LENGTH)) = 5
			- XOR the two. That gives you the byte as grafted together from two entries in the byte array.
		*/
		j := i + offset
		a := content[j]
		b := content[j+1]
		firstPart := a << (8 - (PAYLOAD_LENGTH_LENGTH - VERB_LENGTH))
		secondPart := b >> (8 - (8 - (PAYLOAD_LENGTH_LENGTH - VERB_LENGTH)))
		buf[i] = firstPart ^ secondPart
	}

	return string(buf)
}

// func parseIncomingMessage(conn net.Conn, s *state) (message, error) {
// 	// Messages consist of a verb and a value
// 	// Let's start by parsing the verb. It will take up the first 3 bits of the message
// 	// (room for 8 different verbs).
// 	buffer := make([]byte, 0, 4096)
// 	bytesRead, err := conn.Read(buffer)
// 	if err != nil {
// 		return err
// 	}
// 	if bytesRead < 1 {
// 		return errors.New("Message has the wrong format")
// 	}

// 	// The value is a 64-bit number. Read it as an uint64.

// }
