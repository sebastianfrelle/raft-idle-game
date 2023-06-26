package server

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDecodeMessage(t *testing.T) {
	cases := []struct {
		input []byte
		exp   message
	}{
		{
			input: []byte{
				// Verb is 110, length is 00000110, payload is everything after that.
				0b11000000,
				0b11001100,
				0b00101100,
				0b01001100,
				0b01100000, // last 5 bits of this number is garbage data
			},
			exp: message{
				Verb:    POST,
				Payload: "abd",
			},
		},
	}

	for _, c := range cases {
		act, err := parseMessage(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if !cmp.Equal(*act, c.exp) {
			t.Fatalf("Expected %v; got %v", c.exp, *act)
		}
	}
}

func TestParseVerb(t *testing.T) {
	cases := []struct {
		input []byte
		exp   verb
	}{
		{
			input: []byte{
				// Verb is 110, length is 00000110, payload is everything after that.
				0b11000000,
				0b11001100,
				0b00101100,
				0b01001100,
				0b01100000, // last 5 bits of this number is garbage data
			},
			exp: POST,
		},
	}

	// cases := []struct {
	// 	input []byte
	// 	exp   verb
	// }{
	// 	// TODO: Question: what if we have a message that's longer than 255 bytes? How would we encode the length?
	// 	{input: []byte{0b11000000, 0b11001000}, exp: POST},
	// 	{input: []byte{0b110, 1, 253, 30}, exp: POST},
	// 	{input: []byte{0b110}, exp: POST},
	// }

	for _, c := range cases {
		act, err := parseVerb(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if act != c.exp {
			t.Fatalf("Expected verb %b, got %b", c.exp, act)
		}
	}
}

func TestParsePayloadLength(t *testing.T) {
	cases := []struct {
		input []byte
		exp   uint8
	}{
		{
			input: []byte{
				// Verb is first three bits, length is next 8 bits, payload is everything after that.
				0b11000000,
				0b01001100,
				0b00101100,
				0b01001100,
				0b01100000, // last 5 bits of this number is garbage data
			},
			exp: 0b00000010,
		},
	}

	for _, c := range cases {
		act, err := parsePayloadLength(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if act != c.exp {
			t.Fatalf("Expected %b, got %b", c.exp, act)
		}
	}
}

func TestParsePayload(t *testing.T) {
	cases := []struct {
		content       []byte
		payloadLength uint8
		exp           string
	}{
		{
			content: []byte{
				0b11000000,
				0b01001100,
				0b00101100,
				0b01001100,
				0b01100000,
			},
			payloadLength: 3, // payload length in bytes
			exp:           "abc",
		},
	}

	for _, c := range cases {
		act := parsePayload(c.content, c.payloadLength)
		if act != c.exp {
			t.Fatalf("Expected %s; got %s", c.exp, act)
		}
	}
}

// Message example:
// - 110 - 3-bit number denoting verb
// - 00000011 - 8-bit number denoting length of payload (max length obv 255)
// - 010 - payload. Variable length obv

// Okay, so reading this out as a sequence of bytes yields (if I understand this correctly):
// 0b11000000
// 0b01101000

// So to get the verb out, we have to take the first number and bitshift (8 - 3) places to the right.
// To get the message length out, we have to
// - take the first number
// - bitshift it (8 - 5) = 3 places to the left
// - take the second number
// - bitshift it (8 - 3) = 5 places to the right
// - ...yielding .00000.000
// 					and   00000.011.
// - We can OR or XOR the two to get the length.

// To get the payload out, we
// leftshift the byte to the left by 8 - (PAYLOAD_LENGTH_LENGTH - VERB_LENGTH) = 3
// take the next byte and rightshift it by 8 - (8 - (PAYLOAD_LENGTH_LENGTH - VERB_LENGTH)) = 5
// XOR the two. That gives you the byte at the index.
// 0b11010101
// 0b01010101
// 0b01100000
// XOR each successive number in the byte array

// We need functions for encoding and decoding to and from the protocol
// (really, we're writing a protocol).
